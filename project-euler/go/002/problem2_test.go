package main

import "testing"

func TestProblem2(t *testing.T) {
	expected := 4613732
	got := problem2()
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", expected, got)
	}
}
