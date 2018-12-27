package main

import (
	"bufio"
	"fmt"
	"github.com/ramrunner/coursera-algorithms/collections"
	"os"
	"strings"
)

func isCmd(a string) bool {
	switch a {
	case ".Q", ".AF", ".AB", ".RF", ".RB":
		return true
	}
	return false
}

func (d *dequeCtx) doCmd(cmd string, args ...string) {
	switch cmd {
	case ".Q":
		d.looping = false
	case ".AF":
		for _, s := range args {
			d.dq.AddFirst(s)
		}
	case ".AB":
		for _, s := range args {
			d.dq.AddLast(s)
		}
	case ".RF":
		rs := d.dq.RemoveFirst()
		fmt.Printf("removed front:%s", rs)
	case ".RB":
		rb := d.dq.RemoveLast()
		fmt.Printf("removed back:%s", rb)
	}
	fmt.Printf("stored: %s\n", d.dq.String())
}

type dequeCtx struct {
	dq      collections.StringDeque
	err     error
	looping bool
}

func newDequeCtx() *dequeCtx {
	return &dequeCtx{
		dq:      collections.NewStringDeque(),
		err:     nil,
		looping: true,
	}
}

func (d *dequeCtx) Run() bool {
	return d.looping == true && d.err == nil
}

func (d *dequeCtx) Error() error {
	return d.err
}

func (d *dequeCtx) ProcessLine(parts ...string) {
	if len(parts) < 1 || !isCmd(parts[0]) {
		d.err = fmt.Errorf("malformed line")
		return
	}
	d.doCmd(parts[0], parts[1:]...)
}

func main() {
	fmt.Printf("starting dequeCli:.CMD str1 str2 .. \n.Q to quit .AF add front .AB add back .RF remove front .RB remove back\n")
	inscan := bufio.NewScanner(os.Stdin)
	dq := newDequeCtx()
	for inscan.Scan() {
		sparts := strings.Fields(inscan.Text())
		dq.ProcessLine(sparts...)
		if !dq.Run() {
			break
		}
	}
	if e := dq.Error(); e != nil {
		fmt.Printf("error:%s\n", e)
	}
}
