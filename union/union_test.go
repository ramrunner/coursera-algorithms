package union

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkUnion1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/1000-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, Simple)
		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}

func BenchmarkUnion100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/100-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, Simple)

		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}

func BenchmarkQuickUnion1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/1000-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, QuickUnion)
		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}

func BenchmarkQuickUnion100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/100-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, QuickUnion)

		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}

func BenchmarkWQuickUnion1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/1000-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, QuickUnionWeighted)
		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}

func BenchmarkWQuickUnion100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/100-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, QuickUnionWeighted)

		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}

func BenchmarkPCQuickUnion1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/1000-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, QuickUnionPathCompressed)
		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}

func BenchmarkPCQuickUnion100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open("../datasets/union/100-find")
		if err != nil {
			fmt.Printf("err:%s\n", err)
			return
		}
		ctx := NewUnionCtx(false, f, QuickUnionPathCompressed)

		ConnectAndCheckLoop(ctx)
		f.Close()
	}
}
