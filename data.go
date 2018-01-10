package main

// Should be enough for wide-screen terminals
const HistoryLimit = 1024

// Data holds actual data values, being a buffer for
// UI widges and responsible for flushing/updating/deleting
// old values.
type Data struct {
	cpu     *CircularBuffer
	txBytes *CircularBuffer
	rxBytes *CircularBuffer
}

// NewData inits new data object.
func NewData() *Data {
	return &Data{
		cpu:     NewCircularBuffer(HistoryLimit),
		txBytes: NewCircularBuffer(HistoryLimit),
		rxBytes: NewCircularBuffer(HistoryLimit),
	}
}

// AddCPUValues inserts new CPU measurements into data.
func (d *Data) AddCPUValue(value float64) {
	d.cpu.Add(value)
}

// AddNetworkStats inserts new network stats into data.
func (d *Data) AddNetworkStats(rx, tx float64) {
	d.rxBytes.Add(rx)
	d.txBytes.Add(tx)
}

// CPU returns CPU values.
func (d *Data) CPU() []float64 {
	return d.cpu.Data()
}

// CPU returns CPU values.
func (d *Data) NetworkStats() (rx, tx []float64) {
	return d.rxBytes.Data(), d.txBytes.Data()
}
