package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	serverURL          = "http://localhost:8080"
	localInstallerPath = "/tmp/new_installer.bin"
	currentMarkerFile  = "/home/jeswin-pt8024/myagent/hostzip/.agent_version"
)

func getServerBinaryName() (string, error) {
	resp, err := http.Get(serverURL + "/current-bin")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}

func downloadBinary(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	if err := os.Chmod(dest, 0755); err != nil {
		return err
	}
	return nil
}

func executeBinary(path string) error {
	cmd := exec.Command(path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func readLocalVersion() string {
	b, err := os.ReadFile(currentMarkerFile)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}

func writeLocalVersion(v string) {
	_ = os.WriteFile(currentMarkerFile, []byte(v), 0644)
}

func main() {
	fmt.Println("Checking server binary...")

	serverBin, err := getServerBinaryName()
	if err != nil {
		fmt.Println("Error getting server binary name:", err)
		return
	}

	localVer := readLocalVersion()
	if serverBin == "" {
		fmt.Println("Server returned empty name")
		return
	}

	if serverBin != localVer {
		fmt.Println("New Version available:", serverBin)

		if err := downloadBinary(serverURL+"/download", localInstallerPath); err != nil {
			fmt.Println("Download failed:", err)
			return
		}

		if err := executeBinary(localInstallerPath); err != nil {
			fmt.Println("Executiom failed:", err)
			return
		}

		_ = exec.Command("systemctl", "daemon-reload").Run()
		_ = exec.Command("systemctl", "restart", "logworker.timer").Run()
		_ = exec.Command("systemctl", "restart", "updater.timer").Run()

		writeLocalVersion(serverBin)
		fmt.Println("Update applied:", serverBin)
	} else {
		fmt.Println("No new version")
	}
}
