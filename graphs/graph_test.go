package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestTinyG(t *testing.T) {
	f, err := os.Open("../datasets/graph/tinyG.txt")
	if err != nil {
		fmt.Printf("err:%s\n", err)
		return
	}
	g, err := NewAdjListGraphFromIn(f)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	for i := uint64(0); i < g.V(); i++ {
		for g.Adj(i).Next() {
			v2 := g.Adj(i).Get()
			if v2 != nil {
				fmt.Printf("%d - %d\n", i, v2.(uint64))
			}
		}
	}
}
