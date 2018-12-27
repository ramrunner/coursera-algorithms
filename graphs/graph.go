package graph

import (
	"bufio"
	"fmt"
	"github.com/ramrunner/coursera-algorithms/collections"
	"io"
	"strconv"
	"strings"
)

type Grapher interface {
	AddEdge(uint64, uint64)
	Adj(uint64) collections.Iterable
	V() uint64
	E() uint64
}

type adjlistgraph struct {
	vlist []collections.Bag
	edges uint64
}

func NewAdjListGraph(v uint64) *adjlistgraph {
	ret := &adjlistgraph{
		vlist: make([]collections.Bag, v),
	}
	for i := uint64(0); i < v; i++ {
		ret.vlist[i] = collections.NewSliceBag()
	}
	return ret
}

func (a *adjlistgraph) AddEdge(v, w uint64) {
	a.vlist[v].Add(w)
	a.vlist[w].Add(v)
}

func (a *adjlistgraph) Adj(v uint64) collections.Iterable {
	return a.vlist[v]
}

func (a *adjlistgraph) V() uint64 {
	return uint64(len(a.vlist))
}

func NewAdjListGraphFromIn(in io.Reader) (*adjlistgraph, error) {
	var (
		converr              error
		verts, edges, v1, v2 uint64
		ret                  *adjlistgraph
	)
	curline := 0
	ls := bufio.NewScanner(in)
	for ls.Scan() {
		switch curline {
		case 0:
			verts, converr = strconv.ParseUint(ls.Text(), 10, 64)
		case 1:
			edges, converr = strconv.ParseUint(ls.Text(), 10, 64)
			ret.edges = edges
		default:
			parts := strings.Split(ls.Text(), " ")
			if len(parts) == 2 {
				v1, converr = strconv.ParseUint(parts[0], 10, 64)
				v2, converr = strconv.ParseUint(parts[1], 10, 64)
			} else {
				return nil, fmt.Errorf("malformed line %s", ls.Text())
			}
		}
		if converr != nil {
			return nil, converr
		}
		if curline == 0 {
			ret = NewAdjListGraph(verts)
		} else {
			ret.AddEdge(v1, v2)
		}
		curline++
	}
	return ret, nil
}
