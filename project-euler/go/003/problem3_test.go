package main

import "testing"

func TestProblem3(t *testing.T) {
	expected := 6857
	got := problem3(600851475143)
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", expected, got)
	}
}

func TestGreatestPrimeFactor(t *testing.T) {
	expected := 29
	got := greatest_prime_factor(13195)
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", expected, got)
	}
}

func TestMax(t *testing.T) {
	expected := 3
	got := maxInt([]int{1, 3, 2})
	if got != expected {
		t.Errorf("oh no! expected: %i got: %i", expected, got)
	}
}
