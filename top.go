package main

import (
	"fmt"
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
	fields := parseTopOutput(data)
	if len(fields) < 11 {
		fmt.Println("[ERROR]: wrong top output", fields)
		return nil, ErrParse
	}

	// eithh field is supposed to be CPU value
	cpu, err := strconv.ParseFloat(fields[8], 64)
	if err != nil {
		// top output might be different from system to system, so log
		// this verbosely
		fmt.Println("[ERROR] Parse CPU value:", err)
		fmt.Println("Output:", fields)
		return nil, ErrParse
	}

	return &TopOutput{
		CPU: cpu,
	}, nil
}

func parseTopOutput(data string) []string {
	lines := strings.Split(data, "\r")
	line := strings.Replace(lines[0], "\r", "", -1)
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	return fields
}
