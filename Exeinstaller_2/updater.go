package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	versionURL  = "http://localhost:8080/current-version"
	updateURL   = "http://localhost:8080/update.zip"
	workDir     = "/home/jeswin-pt8024/first-proj"
	versionFile = "/home/jeswin-pt8024/first-proj/.agent_version"
)

func getServerVersion() (string, error) {
	resp, err := http.Get(versionURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func downloadToPath(zipPath string) error {
	resp, err := http.Get(updateURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	zipPtr, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipPtr.Close()

	_, err = io.Copy(zipPtr, resp.Body)
	return err

}

func Unzip(zipFile, destDir string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		destPath := filepath.Join(destDir, f.Name)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(destPath, f.Mode())
			if err != nil {
				return err
			}
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			rc.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	fmt.Println("Entered updater")

	localVersion := "v1"

	data, err := os.ReadFile(versionFile)
	if err == nil {
		tempVer := string(data)
		if tempVer != "" {
			localVersion = tempVer
		}
	} else if os.IsNotExist(err) {
		os.WriteFile(versionFile, []byte(localVersion), 0644)
	}

	serverVersion, err := getServerVersion()
	if err != nil {
		fmt.Println("Error in getting version from server: ", err)
		return
	}

	if localVersion != serverVersion {
		fmt.Println("new version: ", serverVersion, "local version: ", localVersion)

		fmt.Println("Attempting to stop main.service...")

		// Preupgrade stop the service
		cmd := exec.Command("systemctl", "stop", "main.service")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("updater: failed to stop service:", err)
			return
		}
		fmt.Println("main.service stopped successfully.")

		zipPath := workDir + "/update.zip"

		// Download and replace the exe
		err := downloadToPath(zipPath)
		if err != nil {
			fmt.Println("Error in downloading update.zip: ", err)
			exec.Command("systemctl", "start", "main.service").Run()
			return
		}
		defer os.Remove(zipPath)

		err = Unzip(zipPath, workDir+"/hostzip")
		if err != nil {
			fmt.Println("Error in extracting update.zip", err)
			exec.Command("systemctl", "start", "main.service").Run()
			return
		}

		os.WriteFile(versionFile, []byte(serverVersion), 0644)

		// Postupgrade start the srvice
		cmd = exec.Command("systemctl", "start", "main.service")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("updater: failed to restart service:", err)
			return
		}
	}
}
