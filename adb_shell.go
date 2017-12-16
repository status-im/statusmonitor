package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
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

	line := parseTopOutput(out)

	fields := strings.Fields(line)
	if len(fields) < 10 {
		fmt.Println("[ERROR]: wrong top output", fields)
		return 0, ErrParse
	}
	cpu, err := strconv.ParseFloat(fields[9], 64)
	if err != nil {
		// this usually means that app is in background and top
		// omits
		fmt.Println("[ERROR] Parse CPU value:", err)
		fmt.Println("Output:", fields)
		return 0, ErrParse
	}
	return cpu, nil
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

// parseTopOutput parses line from the android shell top's output.
// It contains one line with data and many newlines.
func parseTopOutput(data string) string {
	lines := strings.Split(data, "\r")
	return strings.Replace(lines[0], "\r", "", -1)
}
