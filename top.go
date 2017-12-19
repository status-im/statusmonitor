package main

import (
	"strconv"
	"strings"
)

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
