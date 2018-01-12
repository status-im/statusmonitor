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
	ErrUIDNotFound = errors.New("UID not found")
	ErrCmdFailed   = errors.New("command failed")
	ErrParse       = errors.New("parse failed")
)

// Android implements Source interface for Status app running on
// Android device or emulator via 'adb' tool.
type Android struct {
	pid int64
	// uid is used for parsing /proc/net/xt_qtaguid/stats output
	// TODO: find if there is a way to get stats by PID instead of UID
	uid      int64
	procName string
}

// NewAndroidSource inits new source for Android with given process name.
func NewAndroidSource(procName string) *Android {
	return &Android{
		procName: procName,
	}
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
		if fields[0] == a.procName {
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

func (a *Android) UID() (int64, error) {
	cmd := fmt.Sprintf("dumpsys package %s", a.procName)
	output, err := a.shell(cmd)
	if err != nil {
		return 0, err
	}

	dumpsys, err := NewDumpSysOutput(output)
	if err != nil {
		return 0, ErrUIDNotFound
	}

	return dumpsys.UID, nil
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

func (a *Android) MemStats() (uint64, error) {
	cmd := fmt.Sprintf("dumpsys meminfo -c -s %v", a.pid)
	out, err := a.shell(cmd)
	if err != nil {
		return 0, err
	}

	memInfo, err := NewMemInfoOutput(out)
	if err != nil {
		return 0, err
	}

	return memInfo.UsedMem, nil
}

func (a *Android) Netstats() (int64, int64, error) {
	out, err := a.shell("cat /proc/net/xt_qtaguid/stats")
	if err != nil {
		return 0, 0, err
	}

	netstats, err := NewQTagUIDOutput(out, a.uid)
	if err != nil {
		return 0, 0, err
	}

	return netstats.RxBytes, netstats.TxBytes, nil
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
