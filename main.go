package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gizak/termui"
)

var (
	debug    = flag.Bool("debug", false, "Disable UI and see raw data (debug mode)")
	csvdump  = flag.Bool("csv", false, "Write every point into CSV file [i.e. 20160201_150405.csv]")
	interval = flag.Duration("i", 1*time.Second, "Update interval")
)

func main() {
	flag.Parse()

	pid, err := adbPID()
	if err != nil {
		fmt.Println("Status.im PID not found. Please, make sure that `adb devices` shows your device connected to the computer and Status.im app is launched")
		return
	}
	fmt.Println("Status.im is found on PID", pid)

	if *debug {
		for {
			cpu, err := adbCPU(pid)
			if err != nil {
				fmt.Println("[ERROR]:", err)
				continue
			}
			fmt.Println("CPU:", cpu)
			time.Sleep(*interval)
		}
		return
	}

	// init stuff
	data := NewData()
	var csv *CSVDump
	if *csvdump {
		csv, err = NewCSVDump()
		if err != nil {
			fmt.Println("[ERROR] Can't create csv file, aborting:", err)
			return
		}
	}

	ui := initUI(pid, *interval)
	defer stopUI()

	ui.HandleKeys()

	ui.AddTimer(*interval, func(e termui.Event) {
		cpu, err := adbCPU(pid)
		if err != nil {
			// usually that means app closed or phone disconnected
			stopUI()
			fmt.Println("Disconnected.")
			os.Exit(0)
		}

		// update data
		data.AddCPUValue(cpu)
		if *csvdump {
			csv.Add(cpu)
		}

		ui.UpdateCPU(data.CPU())
		ui.Render()
	})

	ui.Loop()
}
