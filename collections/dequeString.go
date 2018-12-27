package collections

type StringDeque struct {
	dq Dequer
}

func NewStringDeque() StringDeque {
	return StringDeque{
		dq: &Deque{},
	}
}

func (s StringDeque) IsEmpty() bool {
	return s.dq.IsEmpty()
}

func (s StringDeque) AddFirst(a string) {
	s.dq.AddFirst(a)
}

func (s StringDeque) AddLast(a string) {
	s.dq.AddLast(a)
}

func (s StringDeque) RemoveFirst() string {
	return s.dq.RemoveFirst().(string)
}

func (s StringDeque) RemoveLast() string {
	return s.dq.RemoveLast().(string)
}

func (s StringDeque) String() string {
	return s.dq.String()
}
