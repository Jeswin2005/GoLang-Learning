package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	versionURL  = "http://localhost:8080/current-version"
	updateURL   = "http://localhost:8080/update.zip"
	binDir      = "/home/jeswin-pt8024/mini-project/hostzip/bin"
	workDir     = "/home/jeswin-pt8024/mini-project"
	versionFile = "/home/jeswin-pt8024/mini-project/.agent_version"
	serviceName = "main.service"
)

func fetchText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func downloadToPath(url, dst string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}


func extractZipTo(zipPath, dest string) error {
	staging := dest + ".new"

	_ = os.RemoveAll(staging)

	if err := os.MkdirAll(staging, 0755); err != nil {
		return fmt.Errorf("create staging dir: %w", err)
	}

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	// 1) Extract all bin/* entries into the staging directory
	for _, f := range r.File {
		if !strings.HasPrefix(f.Name, "bin/") {
			continue
		}
		rel := strings.TrimPrefix(f.Name, "bin/")
		if rel == "" {
			continue
		}

		fpath := filepath.Join(staging, rel)

		if !strings.HasPrefix(fpath, filepath.Clean(staging)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, f.Mode()); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		in, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			in.Close()
			return err
		}

		if _, err := io.Copy(out, in); err != nil {
			in.Close()
			out.Close()
			return err
		}
		in.Close()
		out.Close()

		_ = os.Chmod(fpath, 0755)
	}

	// 2) Ensure dest dir exists
	if err := os.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("ensure dest dir: %w", err)
	}

	// 3) Move files from staging -> dest atomically using rename
	err = filepath.WalkDir(staging, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(staging, path)
		if err != nil {
			return err
		}
		finalPath := filepath.Join(dest, rel)

		if err := os.MkdirAll(filepath.Dir(finalPath), 0755); err != nil {
			return err
		}

		// os.Rename replaces the destination file atomically.
		if err := os.Rename(path, finalPath); err != nil {
			return fmt.Errorf("rename %s -> %s: %w", path, finalPath, err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// 4) Cleanup staging directory
	_ = os.RemoveAll(staging)

	return nil
}

func main() {
	remote, err := fetchText(versionURL)
	if err != nil {
		fmt.Println("updater: cannot fetch remote version:", err)
		return
	}
	if remote == "" {
		fmt.Println("updater: remote version empty, abort")
		return
	}

	local := "v1"

	if b, err := os.ReadFile(versionFile); err == nil {
		trimmed := strings.TrimSpace(string(b))
		if trimmed != "" {
			local = trimmed
		}
	} else if os.IsNotExist(err) {
		_ = os.WriteFile(versionFile, []byte(local), 0644)
	}


	if remote == local {
		fmt.Println("updater: already at version", local)
		return
	}

	fmt.Println("updater: new version", remote, "available (local:", local, ")")

	zipPath := filepath.Join(workDir, "update.zip")
	if err := downloadToPath(updateURL, zipPath); err != nil {
		fmt.Println("updater: download failed:", err)
		return
	}
	defer os.Remove(zipPath)

	if err := os.MkdirAll(binDir, 0755); err != nil {
		fmt.Println("updater: mkdir failed:", err)
		return
	}
	if err := extractZipTo(zipPath, binDir); err != nil {
		fmt.Println("updater: extract failed:", err)
		return
	}

	if err := os.WriteFile(versionFile, []byte(remote), 0644); err != nil {
		fmt.Println("updater: failed to write version file:", err)
	}

	fmt.Println("updater: restart", serviceName)
	cmd := exec.Command("systemctl", "restart", serviceName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("updater: failed to restart service:", err)
		return
	}

	fmt.Println("updater: update applied to version", remote)
}
