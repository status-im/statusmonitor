package main

import (
	"fmt"
	"strconv"
	"strings"
)

// adbCPU returns CPU value for the given PID via `adb shell` command.
func adbCPU(pid int64) (float64, error) {
	cmd := fmt.Sprintf("top -p %d -n 1 -q -b -o %%CPU", pid)
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

// TopOutput represents data from 'top' command output
// for single process.
type TopOutput struct {
	CPU float64
}

// NewTopOutput creates new TopOutput from raw stdout data.
func NewTopOutput(data string) (*TopOutput, error) {
	line := cleanTopOutput(data)

	cpu, err := strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Println("[ERROR] Parse CPU value:", err)
		return nil, ErrParse
	}

	return &TopOutput{
		CPU: cpu,
	}, nil
}

func cleanTopOutput(data string) string {
	lines := strings.Split(data, "\r")
	line := strings.Replace(lines[0], "\r", "", -1)
	line = strings.Replace(lines[0], "\b", "", -1)
	line = strings.TrimSpace(line)
	return line
}
