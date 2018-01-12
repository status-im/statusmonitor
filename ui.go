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
	rxSparklineIndex
	txSparklineIndex
	numSparklines
)

// UI represents UI layout.
type UI struct {
	Headers    []*termui.Par
	Sparklines *termui.Sparklines

	CPULine        *termui.Sparkline
	RxLine, TxLine *termui.Sparkline

	// fields needed only for UI display
	interval             time.Duration
	maxCPU, maxRx, maxTx float64
}

func initUI(pid int64, interval time.Duration) *UI {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	ui := &UI{
		interval: interval,
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
	line := &ui.Sparklines.Lines[cpuSparklineIndex]
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

// UpdateNetstats updates Netstats widget with new values from data.
func (ui *UI) UpdateNetstats(dataRx, dataTx []float64) {
	intData := make([]int, len(dataRx))

	for i := 0; i < len(dataRx)-1; i++ {
		intData[i] = int(dataRx[i+1] - dataRx[i])
	}

	last := len(dataRx) - 1
	if last > 1 {
		currentRx := dataRx[last] - dataRx[last-1]
		if currentRx > ui.maxRx {
			ui.maxRx = currentRx
		}
		line := &ui.Sparklines.Lines[rxSparklineIndex]
		line.Data = intData
		line.Title = fmt.Sprintf("Network Rx: %v/s, Max: %v/s", byten.Size(int64(currentRx)), byten.Size(int64(ui.maxRx)))
	}

	intData = make([]int, len(dataTx))

	for i := 0; i < len(dataTx)-1; i++ {
		intData[i] = int(dataTx[i+1] - dataTx[i])
	}
	last = len(dataTx) - 1
	if last > 1 {
		currentTx := dataTx[last] - dataTx[last-1]
		if currentTx > ui.maxTx {
			ui.maxTx = currentTx
		}
		line := &ui.Sparklines.Lines[txSparklineIndex]
		line.Data = intData
		line.Title = fmt.Sprintf("Network Tx: %v/s, Max: %v/s", byten.Size(int64(currentTx)), byten.Size(int64(ui.maxTx)))
	}
}

// Render rerenders UI.
func (ui *UI) Render() {
	termui.Body.Align()

	// Update widgets:
	// history time estimation based on new size and interval
	ui.Sparklines.BorderLabel = fmt.Sprintf("Data (last %v)", time.Duration(termui.TermWidth()-2)*ui.interval)
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
			termui.NewCol(12, 0, ui.Sparklines),
		),
	)
}

func (ui *UI) createHeader(pid int64) {
	p := termui.NewPar("")
	p.Height = headerHeight
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = "Monitoring Status.im via adb"
	p.BorderFg = termui.ColorCyan
	p.Text = "press 'q' to exit"

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
	sparklines := make([]termui.Sparkline, numSparklines)

	sCPU := termui.NewSparkline()
	sCPU.Height = (termui.TermHeight() - headerHeight - 3) / cap(sparklines) // - border
	sCPU.LineColor = termui.ColorGreen
	sCPU.Title = "CPU"
	sparklines[cpuSparklineIndex] = sCPU

	ui.CPULine = &sCPU

	sRx := termui.NewSparkline()
	sRx.Height = (termui.TermHeight()-headerHeight-3)/cap(sparklines) - 1 // - border
	sRx.LineColor = termui.ColorGreen
	sRx.Title = "Network Rx"
	sparklines[rxSparklineIndex] = sRx

	sTx := termui.NewSparkline()
	sTx.Height = (termui.TermHeight()-headerHeight-3)/cap(sparklines) - 1 // - border
	sTx.LineColor = termui.ColorGreen
	sTx.Title = "Network Tx"
	sparklines[txSparklineIndex] = sTx

	ui.RxLine = &sRx
	ui.TxLine = &sTx

	// single
	ss := termui.NewSparklines(sparklines...)
	ss.Height = termui.TermHeight() - headerHeight
	ss.Border = true
	ss.BorderLabel = fmt.Sprintf("Data (last %v)", time.Duration(termui.TermWidth()-2)*ui.interval)

	ui.Sparklines = ss
}

func (ui *UI) getUIElements() []termui.Bufferer {
	uiElements := make([]termui.Bufferer, 0)
	for _, h := range ui.Headers {
		uiElements = append(uiElements, h)
	}
	uiElements = append(uiElements, ui.Sparklines)

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
