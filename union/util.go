package union

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	Simple = iota
	QuickUnion
	QuickUnionWeighted
	QuickUnionPathCompressed
)

type UnionCtx struct {
	verbose bool
	r       io.Reader
	variant int
}

func NewUnionCtx(verb bool, r io.Reader, variant int) UnionCtx {
	return UnionCtx{
		verb,
		r,
		variant,
	}
}

func ConnectAndCheckLoop(ctx UnionCtx) {
	scanner := bufio.NewScanner(ctx.r)
	first := true
	var iuni IntUnionFinder
	for scanner.Scan() {
		if first {
			numstr := scanner.Text()
			num, err := strconv.ParseInt(numstr, 10, 32)
			if err != nil {
				fmt.Printf("error parsing set size:%s\n", err)
				return
			}
			switch ctx.variant {
			case Simple:
				iuni = NewIntUnionFind(int(num))
			case QuickUnion:
				iuni = NewIntUnionFindQU(int(num))
			case QuickUnionWeighted:
				iuni = NewIntUnionFindWQU(int(num))
			case QuickUnionPathCompressed:
				iuni = NewIntUnionFindPCQU(int(num))
			default:
				panic("unimplemented UnionFind interface")
			}
			first = false
			continue
		}
		is, err := getIntPair(scanner.Text())
		if err != nil {
			fmt.Printf("error:%s\n", err)
			return
		}
		if con, err := iuni.IsConnected(is[0], is[1]); err == nil {
			if con {
				if ctx.verbose {
					fmt.Printf("Connected!\n")
				}
			} else {
				iuni.Union(is[0], is[1])
			}
		} else {
			fmt.Printf("error:%s", err)
			return
		}
	}

}

func getIntPair(a string) ([]int, error) {
	ret := make([]int, 2)
	parts := strings.Fields(a)
	if len(parts) != 2 {
		return ret, fmt.Errorf("malformed line")
	}
	ia, err1 := strconv.ParseInt(parts[0], 10, 32)
	ib, err2 := strconv.ParseInt(parts[1], 10, 32)
	if err1 != nil || err2 != nil {
		return ret, fmt.Errorf("error parsing int")
	}
	ret[0], ret[1] = int(ia), int(ib)
	return ret, nil
}
