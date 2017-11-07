package main

import (
	"github.com/ramrunner/coursera-algorithms/union"
	"os"
)

func main() {
	ctx := union.NewUnionCtx(true, os.Stdin, uinon.QuickUnion)
	union.ConnectAndCheckLoop(ctx)
}
