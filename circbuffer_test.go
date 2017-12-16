package main

import "testing"

func TestCircularBuffer(t *testing.T) {
	buf := NewCircularBuffer(3)
	got := buf.Data()
	expected := []float64{}
	compareData(t, got, expected)

	buf.Add(5.0)
	buf.Add(10.0)

	got = buf.Data()
	expected = []float64{5.0, 10.0}
	compareData(t, got, expected)

	buf.Add(15.0)
	got = buf.Data()
	expected = []float64{5.0, 10.0, 15.0}
	compareData(t, got, expected)

	buf.Add(20.0)
	got = buf.Data()
	expected = []float64{10.0, 15.0, 20.0}
	compareData(t, got, expected)
}

func compareData(t *testing.T, got, expected []float64) {
	if len(got) != len(expected) {
		t.Fatalf("Expected to get data %d length, got %d", len(expected), len(got))
	}
	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("Expected %v, got %v", expected[i], got[i])
		}
	}
}
