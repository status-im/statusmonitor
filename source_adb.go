package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const StatusIMName = "im.status.ethereum"

var (
	ErrPidNotFound = errors.New("PID not found")
	ErrCmdFailed   = errors.New("command failed")
	ErrParse       = errors.New("parse failed")
)

// Android implements Source interface for Status app running on
// Android device or emulator via 'adb' tool.
type Android struct {
	pid int64
}

func (a *Android) PID() (int64, error) {
	output, err := a.shell("ps -A -w -o CMDLINE,PID")
	if err != nil {
		return 0, err
	}

	lines := strings.Split(output, "\n")
	for i := range lines {
		fields := strings.Fields(lines[i])
		if len(fields) != 2 {
			continue
		}
		if fields[0] == StatusIMName {
			pid, err := strconv.ParseInt(fields[1], 10, 0)
			if err != nil {
				return 0, err
			}
			a.pid = pid
			return pid, nil
		}
	}

	return 0, ErrPidNotFound
}

func (a *Android) CPU() (float64, error) {
	cmd := fmt.Sprintf("top -p %d -n 1 -q -b -o %%CPU", a.pid)
	out, err := a.shell(cmd)
	if err != nil {
		return 0, err
	}

	top, err := NewTopOutput(out)
	if err != nil {
		return 0, err
	}

	return top.CPU, nil
}

func (a *Android) shell(command string) (string, error) {
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
