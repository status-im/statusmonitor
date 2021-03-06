package main

// Data holds actual data values, being a buffer for
// UI widges and responsible for flushing/updating/deleting
// old values.
type Data struct {
	cpu     *CircularBuffer
	mem     *CircularBuffer
	txBytes *CircularBuffer
	rxBytes *CircularBuffer
}

// NewData inits new data object.
func NewData(historyLimit int) *Data {
	return &Data{
		cpu:     NewCircularBuffer(historyLimit),
		mem:     NewCircularBuffer(historyLimit),
		txBytes: NewCircularBuffer(historyLimit),
		rxBytes: NewCircularBuffer(historyLimit),
	}
}

// AddCPUValue inserts new CPU measurements into data.
func (d *Data) AddCPUValue(value float64) {
	d.cpu.Add(value)
}

// AddMemoryStats inserts new memory measurements into data.
func (d *Data) AddMemoryStats(used uint64) {
	d.mem.Add(float64(used))
}

// AddNetworkStats inserts new network stats into data.
func (d *Data) AddNetworkStats(rx, tx int64) {
	d.rxBytes.Add(float64(rx))
	d.txBytes.Add(float64(tx))
}

// CPU returns CPU values.
func (d *Data) CPU() []float64 {
	return d.cpu.Data()
}

// MemoryStats returns memory usage stat values.
func (d *Data) MemoryStats() (used []float64) {
	return d.mem.Data()
}

// NetworkStats returns network stat values.
func (d *Data) NetworkStats() (rx, tx []float64) {
	return d.rxBytes.Data(), d.txBytes.Data()
}

// Clear clears all data.
func (d *Data) Clear() {
	for _, b := range [...]*CircularBuffer{d.cpu, d.mem, d.rxBytes, d.txBytes} {
		b.Reset()
	}
}
