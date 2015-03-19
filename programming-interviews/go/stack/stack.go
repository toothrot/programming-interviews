package stack

// The Element knows its value and what the next element is
type Element struct {
	next  *Element
	Value interface{}
}

// The Stack structure only contains a head element
type Stack struct {
	head *Element
}

// Creates a new stack
func New() *Stack { return new(Stack) }

// Returns the number of items of the stack
func (s *Stack) Len() int {
	if s.head != nil {
		count := 0
		n := s.head

		for {
			if n != nil {
				count++
				n = n.next
			} else {
				break
			}
		}

		return count
	}
	return 0
}

// BUG Not yet implemented
func (s *Stack) Head() *Element {
	return nil
}

// Pushes a new value onto the stack
func (s *Stack) Push(v interface{}) *Element {
	e := new(Element)
	e.Value = v
	e.next = s.head
	s.head = e
	return e
}

// Pops the top value from the stack
func (s *Stack) Pop() interface{} {
	e := s.head
	if e != nil {
		s.head = e.next
		return e.Value
	}
	return nil
}
