package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gizak/termui"
)

var (
	debug    = flag.Bool("debug", false, "Disable UI and see raw data (debug mode)")
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

	data := NewData()
	_ = data
	ui := initUI(pid)
	defer stopUI()

	ui.HandleKeys()

	ui.AddTimer(*interval, func(e termui.Event) {
		cpu, err := adbCPU(pid)
		if err != nil {
			return
		}
		data.AddCPUValue(cpu)

		ui.UpdateCPU(data.CPU())
		ui.Render()
	})

	ui.Loop()
}
