package stack

import "testing"

func TestLenForEmptyStack(t *testing.T) {
	expected := 0
	s := New()
	got := s.Len()

	if got != expected {
		t.Errorf("oh no! expected: %d got: %d", expected, got)
	}
}

func TestLenForOneElement(t *testing.T) {
	expected := 1

	s := New()
	s.Push(1)
	got := s.Len()

	if got != expected {
		t.Errorf("oh no! expected: %d got: %d", expected, got)
	}
}

func TestLenForTwoElements(t *testing.T) {
	expected := 2

	s := New()
	s.Push(1)
	s.Push(1)
	got := s.Len()

	if got != expected {
		t.Errorf("oh no! expected: %d got: %d", expected, got)
	}
}
