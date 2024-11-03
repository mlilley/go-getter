package engine

import (
	"bufio"
	"fmt"
	"os/exec"
)

func getCmd() string {
	return "./vendor/yt-dlp"
}

func Download(url string) error {
	cmd := exec.Command(getCmd(), "--progress", url)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err // todo app error fdsfa
	}

	if err := cmd.Start(); err != nil {
		return err // todo app error
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("OUTPUT:", line)
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err // todo app error
	}

	return nil
}
