package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gizak/termui"
)

// Should be enough for wide-screen terminals
const HistoryLimit = 1024

var (
	debug    = flag.Bool("debug", false, "Disable UI and see raw data (debug mode)")
	csvdump  = flag.Bool("csv", false, "Write every point into CSV file [i.e. 20160201_150405.csv]")
	interval = flag.Duration("i", 1*time.Second, "Update interval")
	source   = flag.String("source", "android", "Data source (android, ios or local)")
	procName = flag.String("proc", "im.status.ethereum", "Process name")
)

func main() {
	flag.Parse()

	src := selectSource(*source, *procName)

	pid, err := src.PID()
	if err != nil {
		fmt.Println("Status.im PID not found. Please, make sure that `adb devices` shows your device connected to the computer and Status.im app is launched")
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status.im is found on PID", pid)

	_, err = src.UID()
	if err != nil {
		fmt.Println("Status.im UID not found. Please, make sure that `adb devices` shows your device connected to the computer and Status.im app is launched")
		fmt.Println("Error:", err)
		return
	}

	if *debug {
		for {
			cpu, err := src.CPU()
			if err != nil {
				fmt.Println("[ERROR]:", err)
				continue
			}
			fmt.Printf("[%v] CPU: %v\n", time.Now().Format(time.RFC3339), cpu)

			usedMem, err := src.MemStats()
			if err != nil {
				fmt.Println("[ERROR]:", err)
				continue
			}
			fmt.Printf("[%v] Used Mem: %v\n", time.Now().Format(time.RFC3339), usedMem)

			rx, tx, err := src.Netstats()
			if err != nil {
				fmt.Println("[ERROR]:", err)
				continue
			}
			fmt.Printf("[%v] Netstats (rx/tx): %v/%v bytes\n", time.Now().Format(time.RFC3339), rx, tx)

			time.Sleep(*interval)
		}
		return
	}

	// init stuff
	data := NewData(HistoryLimit)
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

	// handle key r pressing
	termui.Handle("/sys/kbd/r", func(termui.Event) {
		data.Clear()
	})

	ui.HandleKeys()

	ui.AddTimer(*interval, func(e termui.Event) {
		cpuChan := make(chan float64)
		go func() {
			cpu, err := src.CPU()
			if err != nil {
				// usually that means app closed or phone disconnected
				stopUI()
				fmt.Println("Disconnected.")
				os.Exit(0)
			}

			// update data
			data.AddCPUValue(cpu)
			cpuChan <- cpu
		}()

		usedMemChan := make(chan uint64)
		go func() {
			usedMem, err := src.MemStats()
			if err != nil {
				// usually that means app closed or phone disconnected
				stopUI()
				fmt.Println("Disconnected.")
				os.Exit(0)
			}

			// update data
			data.AddMemoryStats(usedMem)
			usedMemChan <- usedMem
		}()

		// netstats
		netRxChan := make(chan int64)
		netTxChan := make(chan int64)
		go func() {
			rx, tx, err := src.Netstats()
			if err != nil {
				// usually that means app closed or phone disconnected
				stopUI()
				fmt.Println("Disconnected.")
				os.Exit(0)
			}

			data.AddNetworkStats(rx, tx)
			netRxChan <- rx
			netTxChan <- tx
		}()

		cpu := <-cpuChan
		usedMem := <-usedMemChan
		rx := <-netRxChan
		tx := <-netTxChan

		// csv
		if *csvdump {
			csv.Add(cpu, usedMem, rx, tx)
		}

		ui.UpdateCPU(data.CPU())
		ui.UpdateMemoryStats(data.MemoryStats())
		ui.UpdateNetstats(data.NetworkStats())
		ui.Render()
	})

	ui.Loop()
}

func selectSource(source, procName string) Source {
	switch source {
	case "android":
		return NewAndroidSource(procName)
	case "ios":
		log.Fatal("iOS source not implemented yet")
	case "local":
		return &Local{}
	}
	log.Fatal("Incorrect source")
	return nil
}
