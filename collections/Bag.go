package collections

type slicebag struct {
	elems []interface{}
	iterr error
	itpos int
}

type Bag interface {
	Iterable
	Add(a interface{})
	IsEmpty() bool
	Size() int
}

func NewSliceBag() *slicebag {
	return &slicebag{
		elems: make([]interface{}, 0),
	}
}

func (s *slicebag) Add(a interface{}) {
	s.elems = append(s.elems, a)
}

func (s *slicebag) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *slicebag) Size() int {
	return len(s.elems)
}

//make it iterable with Next(), Get() and Err()
func (s *slicebag) Next() bool {
	if s.itpos < len(s.elems) && s.iterr == nil {
		s.itpos++
		return true
	}
	return false
}

func (s *slicebag) Get() interface{} {
	if s.itpos < len(s.elems) {
		return s.elems[s.itpos]
	}
	return nil
}

func (s *slicebag) Err() error {
	return s.iterr
}
