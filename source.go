package main

// Source represents data source for monitoring.
type Source interface {
	PID() (int64, error)
	CPU() (float64, error)
}
