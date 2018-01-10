package main

import (
	"fmt"
	"strconv"
	"strings"
)

// QTagUIDOutput represents data from '/proc/net/xt_qtaguid/stats' kernel output.
type QTagUIDOutput struct {
	RxBytes int64
	TxBytes int64
}

// NewQTagUIDOutput creates new TopOutput from raw stdout data.
func NewQTagUIDOutput(data string, uid int64) (*QTagUIDOutput, error) {
	lines := strings.Split(data, "\n")

	var totalRx, totalTx int64
	uidStr := fmt.Sprintf("%d", uid)
	for i := range lines {
		fields := strings.Fields(lines[i])
		if len(fields) < 8 {
			continue
		}

		if strings.TrimSpace(fields[3]) != uidStr {
			continue
		}

		// get 6 and 8 fields (rx and tx accordingly)
		rx, err := strconv.ParseInt(fields[5], 10, 0)
		if err != nil {
			continue
		}
		tx, err := strconv.ParseInt(fields[7], 10, 0)
		if err != nil {
			continue
		}

		// sum up values
		totalRx += rx
		totalTx += tx
	}

	return &QTagUIDOutput{
		RxBytes: totalRx,
		TxBytes: totalTx,
	}, nil
}
