package main

import "testing"

func TestProblem(t *testing.T) {
	expected := 232792560
	got := problem()
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", expected, got)
	}
}
