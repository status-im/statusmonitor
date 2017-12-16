package main

import (
	"fmt"
	"time"

	"github.com/gizak/termui"
)

// UI represents UI layout.
type UI struct {
	Header     *termui.Par
	Sparklines *termui.Sparklines

	CPULine *termui.Sparkline
}

func initUI(pid int64) *UI {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	ui := &UI{}
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

// Render rerenders UI.
func (ui *UI) Render() {
	termui.Render(ui.Header, ui.Sparklines)
}

// Align recalculates layout and aligns widgets.
func (ui *UI) Align() {
	termui.Body.Align()
}

func (ui *UI) createLayout() {
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, ui.Header),
		),
		termui.NewRow(
			termui.NewCol(12, 0, ui.Sparklines),
		),
	)
}

func (ui *UI) createHeader(pid int64) {
	p := termui.NewPar("Press `q` to exit")
	p.Height = 3
	p.TextFgColor = termui.ColorWhite
	p.BorderLabel = fmt.Sprintf("Monitoring Status.im app on PID %d", pid)
	p.BorderFg = termui.ColorCyan

	ui.Header = p
}

func (ui *UI) createSparklines() {
	s := termui.NewSparkline()
	s.Height = termui.TermHeight() - ui.Header.Height - 3 // - border
	s.Title = "Status.im CPU"
	s.LineColor = termui.ColorGreen

	ui.CPULine = &s

	// single
	ss := termui.NewSparklines(s)
	ss.Height = termui.TermHeight() - ui.Header.Height
	ss.Border = true

	ui.Sparklines = ss
}

func (ui *UI) HandleKeys() {
	// handle key q pressing
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
}

// AddTimer adds handler for repeatable functions that interact with UI.
func (ui *UI) AddTimer(d time.Duration, fn func(e termui.Event)) {
	durationStr := fmt.Sprintf("/timer/%v", d)
	termui.Handle(durationStr, fn)
}

// Loop starts event loop and blocks.
func (ui *UI) Loop() {
	termui.Loop()
}
