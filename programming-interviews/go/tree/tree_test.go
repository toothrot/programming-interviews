package tree

import "testing"

func TestHeightOfSingleNode(t *testing.T) {
	expected := 1
	n := New(1)
	got := n.Height()

	if got != expected {
		t.Errorf("oh no! expected: %d got: %d", expected, got)
	}
}

func TestHeight(t *testing.T) {
	expected := 3
	n := New(5)
	n.Insert(4)
	n.Insert(3)
	got := n.Height()

	if got != expected {
		t.Errorf("oh no! expected: %d got: %d", expected, got)
	}

	expected = 2
	n = New(4)
	n.Insert(5)
	n.Insert(3)

	got = n.Height()

	if got != expected {
		t.Errorf("oh no! expected: %d got: %d", expected, got)
	}

	n.Insert(2)
	expected = 3

	got = n.Height()

	if got != expected {
		t.Errorf("oh no! expected: %d got: %d", expected, got)
	}
}
