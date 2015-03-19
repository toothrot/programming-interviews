package stack

type Element struct {
	next  *Element
	Value interface{}
}

type Stack struct {
	head *Element
}

func New() *Stack { return new(Stack) }

func (s *Stack) Len() int {
	if s.head == nil {
		return 0
	} else {
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
}

func (s *Stack) Head() *Element {
	return nil
}

func (s *Stack) Push(interface{}) *Element {
	e := new(Element)
	e.next = s.head
	s.head = e
	return e
}
