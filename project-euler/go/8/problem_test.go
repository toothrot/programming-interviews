package main

import "testing"

func TestProblem(t *testing.T) {
	expected := 23514624000
	got := problem()
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", expected, got)
	}
}
