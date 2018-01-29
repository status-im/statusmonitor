package main

// Local is a dummy implementation of the Source interface.
type Local struct {
}

func (l *Local) PID() (int64, error) {
	return 0, nil
}

func (l *Local) UID() (int64, error) {
	return 0, nil
}

func (l *Local) CPU() (float64, error) {
	return 0, nil
}

func (l *Local) MemStats() (uint64, error) {
	return 0, nil
}

func (l *Local) Netstats() (int64, int64, error) {
	return 0, 0, nil
}
