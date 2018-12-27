package collections

import (
	"errors"
	"math/rand"
)

var (
	enoitem = errors.New("no such item")
)

type RandomizedQueuer interface {
	IsEmpty() bool
	Size() int
	Enqueue(a interface{})
	Dequeue() interface{}
	Sample() interface{}
}

type RandomizedQueue struct {
	stack []interface{}
	N     int
}

func NewRandomizedQueue() *RandomizedQueue {
	return &RandomizedQueue{
		stack: make([]interface{}, 1),
	}
}

func (r *RandomizedQueue) IsEmpty() bool {
	return r.N == 0
}

func (r *RandomizedQueue) Size() int {
	return r.N
}

func (r *RandomizedQueue) resize(sz int) {
	stackcp := make([]interface{}, sz)
	for i := 0; i < r.N; i++ {
		stackcp[i] = r.stack[i]
	}
	r.stack = stackcp
}

func (r *RandomizedQueue) Enqueue(a interface{}) {
	if r.N == len(r.stack) {
		r.resize(2 * len(r.stack))
	}
	r.stack[r.N] = a
	r.N++
}

func (r *RandomizedQueue) Dequeue() interface{} {
	if r.N <= 0 {
		return nil
	}
	removei := rand.Intn(r.N)
	temp := r.stack[removei]
	r.stack[removei] = r.stack[r.N-1]
	r.N = r.N - 1
	r.stack[r.N] = nil
	//free memory if we are on 1/4th of the allocated space
	if r.N > 0 && r.N == len(r.stack)/4 {
		r.resize(len(r.stack) / 2)
	}
	return temp
}

func (r *RandomizedQueue) Sample() interface{} {
	return r.stack[rand.Intn(r.N)]
}
