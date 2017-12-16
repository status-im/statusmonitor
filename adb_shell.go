package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	ErrCmdFailed = errors.New("command failed")
	ErrParse     = errors.New("parse failed")
)

// adbShell calls custom command via 'adb shell` and returns it's stdout output.
func adbShell(command string) (string, error) {
	cmdWords := strings.Fields(command)
	args := []string{"shell"}
	args = append(args, cmdWords...)
	cmd := exec.Command("adb", args...)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		fmt.Println("[ERROR] adbShell:", err)
		return "", ErrCmdFailed
	}

	return buf.String(), nil
}
