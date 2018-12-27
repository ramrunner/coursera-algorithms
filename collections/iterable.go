package collections

type Iterable interface {
	Next() bool
	Get() interface{}
	Err() error
}
