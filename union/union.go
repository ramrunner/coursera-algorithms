package union

import (
	"fmt"
)

type IntUnionFinder interface {
	IsConnected(int, int) (bool, error)
	Union(int, int) error
}

type IntUnionFind struct {
	arr []int
}

func NewIntUnionFind(sz int) *IntUnionFind {
	arr := make([]int, sz)
	for i := range arr {
		arr[i] = i
	}
	return &IntUnionFind{
		arr: arr,
	}
}

func (i *IntUnionFind) IsConnected(a, b int) (bool, error) {
	if a >= len(i.arr) || b >= len(i.arr) {
		return false, fmt.Errorf("larger number than set supports")
	}
	return i.arr[a] == i.arr[b], nil

}

func (i *IntUnionFind) Union(a, b int) error {
	if a >= len(i.arr) || b >= len(i.arr) {
		return fmt.Errorf("larger number than set supports")
	}
	changing := i.arr[a]
	to := i.arr[b]
	for ind := range i.arr {
		if i.arr[ind] == changing {
			i.arr[ind] = to
		}
	}
	return nil
}

type IntUnionFindQU struct {
	arr []int
}

func NewIntUnionFindQU(sz int) *IntUnionFindQU {
	arr := make([]int, sz)
	for i := range arr {
		arr[i] = i
	}
	return &IntUnionFindQU{
		arr: arr,
	}
}

func (i *IntUnionFindQU) root(a int) int {
	for a != i.arr[a] {
		a = i.arr[a]
	}
	return a
}

func (i *IntUnionFindQU) IsConnected(a, b int) (bool, error) {
	if a >= len(i.arr) || b >= len(i.arr) {
		return false, fmt.Errorf("larger number than set supports")
	}
	return i.root(a) == i.root(b), nil
}

func (i *IntUnionFindQU) Union(a, b int) error {
	if a >= len(i.arr) || b >= len(i.arr) {
		return fmt.Errorf("larger number than set supports")
	}
	p := i.root(a)
	j := i.root(b)
	i.arr[p] = j
	return nil
}

type IntUnionFindWQU struct {
	IntUnionFindQU
	sz []int
}

func NewIntUnionFindWQU(sz int) *IntUnionFindWQU {
	arr, siz := make([]int, sz), make([]int, sz)
	for i := range arr {
		arr[i], siz[i] = i, i
	}
	ufqi := NewIntUnionFindQU(sz)
	return &IntUnionFindWQU{
		*ufqi,
		siz,
	}
}

func (i *IntUnionFindWQU) Union(a, b int) error {
	if a >= len(i.arr) || b >= len(i.arr) {
		return fmt.Errorf("larger number than set supports")
	}
	p := i.root(a)
	j := i.root(b)
	if p == j {
		return nil
	}
	if i.sz[p] < i.sz[j] {
		i.arr[p] = j
		i.sz[j] += i.sz[p]
	} else {
		i.arr[j] = p
		i.sz[p] += i.sz[j]
	}
	return nil
}

type IntUnionFindPCQU struct {
	IntUnionFindQU
}

func NewIntUnionFindPCQU(sz int) *IntUnionFindPCQU {
	ufqi := NewIntUnionFindQU(sz)
	return &IntUnionFindPCQU{
		*ufqi,
	}
}

func (i *IntUnionFindPCQU) root(a int) int {
	for a != i.arr[a] {
		i.arr[a] = i.arr[i.arr[a]] //compress the path to half
		a = i.arr[a]
	}
	return a
}
