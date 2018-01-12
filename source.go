package main

// Source represents data source for monitoring.
type Source interface {
	PID() (int64, error)
	UID() (int64, error)
	CPU() (float64, error)
	MemStats() (uint64, error)
	Netstats() (int64, int64, error)
}
