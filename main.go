package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gizak/termui"
)

var (
	debug = flag.Bool("noui", false, "Disable UI and see raw data (debug mode)")
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
			time.Sleep(1 * time.Second)
		}
		return
	}

	ui := initUI(pid)
	defer stopUI()

	ui.HandleKeys()

	ui.AddTimer(1*time.Second, func(e termui.Event) {
		cpu, err := adbCPU(pid)
		if err != nil {
			return
		}
		_ = cpu

		updateData(ui.Sparklines)
		ui.Render()
	})

	ui.Loop()
}

func updateData(s *termui.Sparklines) {
	data := []int{2, 2, 2, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 111, 2, 3, 23, 3, 22, 77, 8}
	s.Lines[0].Data = data
}
