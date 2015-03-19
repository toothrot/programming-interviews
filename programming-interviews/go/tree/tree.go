// Binary Search Tree
package tree

func maxInt(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

type Tree struct {
	left  *Tree
	right *Tree
	Value int
}

func New(v int) *Tree {
	t := new(Tree)
	t.Value = v
	return t
}

func (t *Tree) Height() int {
	height := 1
	maxLeft := 0
	maxRight := 0

	if t.left != nil {
		maxLeft = t.left.Height()
	}

	if t.right != nil {
		maxRight = t.right.Height()
	}

	return height + maxInt(maxLeft, maxRight)
}

func (t *Tree) insertLeft(v int) {
	if t.left != nil {
		t.left.Insert(v)
	} else {
		t.left = New(v)
	}
}

func (t *Tree) insertRight(v int) {
	if t.right != nil {
		t.right.Insert(v)
	} else {
		t.right = New(v)
	}
}

func (t *Tree) Insert(v int) {
	if v <= t.Value {
		t.insertLeft(v)
	} else {
		t.insertRight(v)
	}
}
