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

func TestPop(t *testing.T) {
	expected := "abc"

	s := New()
	s.Push("abc")
	got := s.Pop()

	if got != expected {
		t.Errorf("oh no! expected: %s got: %s", expected, got)
	}
}

func TestPopTwoThings(t *testing.T) {
	s := New()
	s.Push("abc")
	s.Push("def")

	expected := "def"
	got := s.Pop()

	if got != expected {
		t.Errorf("oh no! expected: %s got: %s", expected, got)
	}

	expected = "abc"
	got = s.Pop()

	if got != expected {
		t.Errorf("oh no! expected: %s got: %s", expected, got)
	}
}

func TestPopTooManyThings(t *testing.T) {
	s := New()
	s.Push("abc")
	s.Push("def")

	s.Pop()
	s.Pop()
	got := s.Pop()

	if got != nil {
		t.Errorf("oh no! expected: %s got: %s", nil, got)
	}
}
