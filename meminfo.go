package main

import (
	"strconv"
	"strings"
)

// MemInfoOutput represents data from 'dumpsys meminfo' command output
// for single process.
type MemInfoOutput struct {
	UsedMem uint64
}

// NewMemInfoOutput creates new MemInfoOutput from raw stdout data.
func NewMemInfoOutput(data string) (*MemInfoOutput, error) {
	line := cleanMemInfoOutput(data)

	usedMem, err := strconv.ParseUint(line, 10, 64)
	if err != nil {
		return nil, ErrParse
	}

	return &MemInfoOutput{
		UsedMem: usedMem * 1024,
	}, nil
}

func cleanMemInfoOutput(data string) string {
	if index := strings.Index(data, "TOTAL:"); index >= 0 {
		data = strings.TrimSpace(data[index+6:])
		index := strings.Index(data, " ")
		data = data[:index]
		return data
	}
	return ""
}
