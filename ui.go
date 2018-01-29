package main

import (
	"fmt"
	"time"

	"github.com/gizak/termui"
	"github.com/pyk/byten"
)

const headerHeight = 3
const (
	cpuSparklineIndex = iota
	usedMemSparklineIndex
)
const (
	rxSparklineIndex = iota
	txSparklineIndex
)

// UI represents UI layout.
type UI struct {
	Headers                         []*termui.Par
	SparklinesLeft, SparklinesRight *termui.Sparklines

	CPULine        *termui.Sparkline
	UsedMemLine    *termui.Sparkline
	RxLine, TxLine *termui.Sparkline

	// fields needed only for UI display
	interval                                     time.Duration
	maxCPU, minUsedMem, maxUsedMem, maxRx, maxTx float64
}

func initUI(pid int64, interval time.Duration) *UI {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	ui := &UI{
		interval:   interval,
		minUsedMem: -1.,
	}
	ui.createHeader(pid)
	ui.createSparklines()
	ui.createLayout()

	ui.Align()
	ui.Render()

	return ui
}

func stopUI() {
	termui.Close()
}

// UpdateCPU updates CPU widget with new values from data.
func (ui *UI) UpdateCPU(data []float64) {
	intData := make([]int, len(data))

	// multiply by 100, to account for 2 decimals after the point
	for i := range data {
		intData[i] = int(data[i] * 100)
	}
	line := &ui.SparklinesLeft.Lines[cpuSparklineIndex]
	line.Data = intData

	if len(data) == 0 {
		return
	}

	currentCPU := data[len(data)-1]
	if currentCPU > ui.maxCPU {
		ui.maxCPU = currentCPU
	}
	line.Title = fmt.Sprintf("CPU: %.2f%%, Max: %.2f%%", currentCPU, ui.maxCPU)
}

// UpdateMemoryStats updates memory usage widget with new values from data.
func (ui *UI) UpdateMemoryStats(data []float64) {
	intData := make([]int, len(data))

	// multiply by 100, to account for 2 decimals after the point
	for i := range data {
		intData[i] = int(data[i] * 100)
	}
	line := &ui.SparklinesLeft.Lines[usedMemSparklineIndex]
	line.Data = intData

	if len(data) == 0 {
		return
	}

	currentUsedMem := data[len(data)-1]
	if ui.minUsedMem < 0. || currentUsedMem < ui.minUsedMem {
		ui.minUsedMem = currentUsedMem
	}
	if currentUsedMem > ui.maxUsedMem {
		ui.maxUsedMem = currentUsedMem
	}
	line.Title = fmt.Sprintf("Used mem: %v, Min: %v, Max: %v", memSize(currentUsedMem), memSize(ui.minUsedMem), memSize(ui.maxUsedMem))
}

// UpdateNetstats updates Netstats widget with new values from data.
func (ui *UI) UpdateNetstats(rxData, txData []float64) {
	ui.updateNetstats(rxSparklineIndex, rxData)
	ui.updateNetstats(txSparklineIndex, txData)
}

func (ui *UI) updateNetstats(sparklineIndex int, data []float64) {
	intData := make([]int, len(data))

	for i := 0; i < len(data)-1; i++ {
		intData[i] = int(data[i+1] - data[i])
	}

	last := len(data) - 1
	if last > 1 {
		var direction string
		var max *float64
		switch sparklineIndex {
		case rxSparklineIndex:
			max = &ui.maxRx
			direction = "Rx"
		case txSparklineIndex:
			max = &ui.maxTx
			direction = "Tx"
		}

		total := data[last]
		current := total - data[last-1]
		if current > *max {
			*max = current
		}
		line := &ui.SparklinesRight.Lines[sparklineIndex]
		line.Data = intData
		line.Title = fmt.Sprintf("Network %s: %v/s, Max: %v/s (total: %v)", direction, memSize(current), memSize(*max), memSize(total))
	}
}

// Render rerenders UI.
func (ui *UI) Render() {
	termui.Body.Align()

	// Update widgets:
	// history time estimation based on new size and interval
	ui.SparklinesLeft.BorderLabel = fmt.Sprintf("Data (last %v)", time.Duration(termui.TermWidth()-2)*ui.interval)
	// time in header
	ui.Headers[3].Text = fmt.Sprintf("%v", time.Now().Format("15:04:05"))

	termui.Render(ui.getUIElements()...)
}

// Align recalculates layout and aligns widgets.
func (ui *UI) Align() {
	termui.Body.Align()
}

func (ui *UI) createLayout() {
	if len(ui.Headers) != 4 {
		panic("update headers code")
	}
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(6, 0, ui.Headers[0]),
			termui.NewCol(2, 0, ui.Headers[1]),
			termui.NewCol(2, 0, ui.Headers[2]),
			termui.NewCol(2, 0, ui.Headers[3]),
		),
		termui.NewRow(
			termui.NewCol(6, 0, ui.SparklinesLeft),
			termui.NewCol(6, 0, ui.SparklinesRight),
		),
	)
}

func (ui *UI) createHeader(pid int64) {
	p := termui.NewPar("")
	p.Height = headerHeight
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = "Monitoring Status.im via adb"
	p.BorderFg = termui.ColorCyan
	p.Text = "press 'q' to exit, 'r' to reset"

	p1 := termui.NewPar("")
	p1.Height = headerHeight
	p1.TextFgColor = termui.ColorWhite
	p1.BorderLabel = "PID"
	p1.BorderFg = termui.ColorCyan
	p1.Text = fmt.Sprintf("%d", pid)

	p2 := termui.NewPar("")
	p2.Height = headerHeight
	p2.TextFgColor = termui.ColorWhite
	p2.BorderLabel = "Polling interval"
	p2.BorderFg = termui.ColorCyan
	p2.Text = fmt.Sprintf("%v", ui.interval)

	p3 := termui.NewPar("")
	p3.Height = headerHeight
	p3.TextFgColor = termui.ColorYellow
	p3.BorderLabel = "Time"
	p3.BorderFg = termui.ColorCyan
	p3.Text = fmt.Sprintf("%v", time.Now().Format("15:04:05"))

	ui.Headers = []*termui.Par{p, p1, p2, p3}
}

func (ui *UI) createSparklines() {
	sparklines1 := make([]termui.Sparkline, 2)

	sCPU := termui.NewSparkline()
	sCPU.Height = (termui.TermHeight() - headerHeight - 3) / 2
	sCPU.LineColor = termui.ColorGreen
	sCPU.Title = "CPU"
	sparklines1[cpuSparklineIndex] = sCPU

	ui.CPULine = &sCPU

	sUsedMem := termui.NewSparkline()
	sUsedMem.Height = (termui.TermHeight()-headerHeight-3)/2 - 1
	sUsedMem.LineColor = termui.ColorGreen
	sUsedMem.Title = "Used Mem"
	sparklines1[usedMemSparklineIndex] = sUsedMem

	ui.UsedMemLine = &sUsedMem

	sparklines2 := make([]termui.Sparkline, 2)
	sRx := termui.NewSparkline()
	sRx.Height = (termui.TermHeight() - headerHeight - 3) / 2
	sRx.LineColor = termui.ColorGreen
	sRx.Title = "Network Rx"
	sparklines2[rxSparklineIndex] = sRx

	sTx := termui.NewSparkline()
	sTx.Height = (termui.TermHeight() - headerHeight - 3) / 2
	sTx.LineColor = termui.ColorGreen
	sTx.Title = "Network Tx"
	sparklines2[txSparklineIndex] = sTx

	ui.RxLine = &sRx
	ui.TxLine = &sTx

	sleft := termui.NewSparklines(sparklines1...)
	sleft.Height = termui.TermHeight() - headerHeight
	sleft.Border = true
	sleft.BorderLabel = fmt.Sprintf("Data (last %v)", time.Duration(termui.TermWidth()-2)*ui.interval)

	sright := termui.NewSparklines(sparklines2...)
	sright.Height = termui.TermHeight() - headerHeight
	sright.Border = true
	sright.BorderLabel = fmt.Sprintf("Data (last %v)", time.Duration(termui.TermWidth()-2)*ui.interval)

	ui.SparklinesLeft = sleft
	ui.SparklinesRight = sright
}

func (ui *UI) getUIElements() []termui.Bufferer {
	uiElements := make([]termui.Bufferer, 0)
	for _, h := range ui.Headers {
		uiElements = append(uiElements, h)
	}
	uiElements = append(uiElements, ui.SparklinesLeft, ui.SparklinesRight)

	return uiElements
}

func (ui *UI) HandleKeys() {
	// handle key q pressing
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
}

// AddTimer adds handler for repeatable functions that interact with UI.
func (ui *UI) AddTimer(d time.Duration, fn func(e termui.Event)) {
	durationStr := fmt.Sprintf("/timer/%s", d)
	termui.Handle(durationStr, fn)
}

// Loop starts event loop and blocks.
func (ui *UI) Loop() {
	termui.Loop()
}

// memSize is a wrapper around value.
func memSize(value float64) string {
	return byten.Size(int64(value))
}
