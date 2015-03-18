package main

import "testing"

func TestProblem4(t *testing.T) {
	expected := 906609
	got := problem4()
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", expected, got)
	}
}

func TestIsPalendrome(t *testing.T) {
	got := isPalendrome(101)
	if got != true {
		t.Errorf("oh no! expected: %i got: %i", true, got)
	}

	got = isPalendrome(1)
	if got != true {
		t.Errorf("oh no! expected: %i got: %i", true, got)
	}

	got = isPalendrome(11)
	if got != true {
		t.Errorf("oh no! expected: %i got: %i", true, got)
	}

	got = isPalendrome(12)
	if got != false {
		t.Errorf("oh no! expected: %i got: %i", true, got)
	}
}
