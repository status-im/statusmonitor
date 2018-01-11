package main

import (
	"fmt"
	"os"
	"time"
)

// CSVDump is responsible for dumping CPU data in CSV format.
type CSVDump struct {
	file string
}

// New CSVDump creates new CSVDump object.
func NewCSVDump() (*CSVDump, error) {
	now := time.Now()
	filename := now.Format("20060102_150405.csv")

	fd, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	fd.WriteString("timestamp,cpu,rx,tx\n")
	fd.Close()

	return &CSVDump{
		file: filename,
	}, nil
}

// Adds adds new CPU value to the CSV dump.
func (c *CSVDump) Add(cpu float64, rx, tx int64) {
	fd, err := os.OpenFile(c.file, os.O_APPEND|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		// we just created this file, it should not dissapear :D
		panic(err)
	}

	now := time.Now()

	fd.WriteString(fmt.Sprintf("%d,%f,%d,%d\n", now.Unix(), cpu, rx, tx))
	fd.Close()
}
