package main

// Should be enough for wide-screen terminals
const HistoryLimit = 1024

// Data holds actual data values, being a buffer for
// UI widges and responsible for flushing/updating/deleting
// old values.
type Data struct {
	cpu *CircularBuffer
}

// NewData inits new data object.
func NewData() *Data {
	return &Data{
		cpu: NewCircularBuffer(HistoryLimit),
	}
}

func (d *Data) AddCPUValue(value float64) {
	d.cpu.Add(value)
}

func (d *Data) CPU(n int) []float64 {
	return d.cpu.Data()
}
