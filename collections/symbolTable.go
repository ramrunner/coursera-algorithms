// +build unfinished

package collections

type SymbolTabler interface {
	Put(Keyer, interface{})
	Get(Keyer) (interface{}, bool)
	Delete(Keyer) error
	Contains(Keyer) bool
	IsEmpty() bool
	Size() int
	Next() (Keyer, bool)
	Error() error
}

type Keyer interface {
	Equals(Keyer) bool
}

type strBsST struct {
	keys []string
	vals []interface{}
}

func (s *strBsST) Get(k Keyer) (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	//rank := rank(k)
	if i < len(s.vals)-1 && s.keys[i].Equals(k) {
		return s.vals[i], true
	}
	return nil, false
}

func (s *strBsST) rank(elems int, k Keyer) {
	lo, hi := 0, elems
	for lo <= hi {
		mid := (hi + lo) >> 1
		if ran := k.CompareTo(s.keys[mid]); ran < 0 {
			hi = mid - 1
		} else if ran > 0 {
			lo = mid + 1
		} else {
			return mid
		}
		return lo
	}
}

func (s *strBsST) IsEmpty() bool {
	return len(keys) == 0
}
