package main

type Local struct {
}

func (l *Local) PID() (int64, error) {
	return 0, nil
}

func (l *Local) CPU() (float64, error) {
	return 0, nil
}
