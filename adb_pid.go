package main

import (
	"errors"
	"strconv"
	"strings"
)

const StatusIMName = "im.status.ethereum"

var ErrPidNotFound = errors.New("PID not found")

// adbPID finds PID of Status.im running process on the device.
func adbPID() (int64, error) {
	output, err := adbShell("ps -A -w -o CMDLINE,PID")
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
			return pid, nil
		}
	}

	return 0, ErrPidNotFound
}
