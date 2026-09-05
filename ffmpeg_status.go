package main

import (
	"os/exec"
)

func CheckFFmpeg() bool {

	cmd := exec.Command(
		ffmpegPath(),
		"-version",
	)

	err := cmd.Run()

	return err == nil
}
