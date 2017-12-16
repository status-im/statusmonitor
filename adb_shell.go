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

// adbCPU returns CPU value for the given PID via `adb shell` command.
func adbCPU(pid int64) (float64, error) {
	cmd := fmt.Sprintf("top -p %d -n 1 -q", pid)
	out, err := adbShell(cmd)
	if err != nil {
		return 0, err
	}

	top, err := NewTopOutput(out)
	if err != nil {
		return 0, err
	}

	return top.CPU, nil
}

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
