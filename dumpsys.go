package main

import (
	"errors"
	"strconv"
	"strings"
)

// ErrInvalidOutput represents parsing error.
var ErrInvalidOutput = errors.New("invalid output")

// DumpSysOutput represents data from 'dumpsys package' command output.
type DumpSysOutput struct {
	UID int64
}

// NewDumpSysOutput creates new TopOutput from raw stdout data.
func NewDumpSysOutput(data string) (*DumpSysOutput, error) {
	lines := strings.Split(data, "\n")

	var uid int64
	for i := range lines {
		if !strings.Contains(lines[i], "userId") {
			continue
		}

		str := strings.TrimSpace(lines[i])
		fields := strings.Split(str, "=")
		// should be "userID=XXX" format
		if len(fields) != 2 {
			return nil, ErrInvalidOutput
		}

		value, err := strconv.ParseInt(fields[1], 10, 0)
		if err != nil {
			return nil, ErrInvalidOutput
		}

		uid = value
		break
	}

	return &DumpSysOutput{
		UID: uid,
	}, nil
}
