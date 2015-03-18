package main

import "testing"

func TestProblem1(t *testing.T) {
	expected := 233168
	got := problem1()
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", got, expected)
	}

}
