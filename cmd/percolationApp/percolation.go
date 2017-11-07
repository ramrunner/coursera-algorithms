package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ramrunner/coursera-algorithms/union"
	"net/http"
	"strconv"
	"sync"
)

const (
	boxsz      = 5
	boxspacing = 1
)

type percappctx struct {
	boards map[string]*percboard
	router *mux.Router
}

func newPercAppCtx() *percappctx {
	return &percappctx{
		boards: make(map[string]*percboard),
		router: mux.NewRouter(),
	}
}

type percboard struct {
	name string
	sz   int
	con  map[int]bool
	uf   union.IntUnionFinder
	uf1  union.IntUnionFinder
	mut  *sync.Mutex
}

func newPercBoard(name string, sz int) *percboard {
	return &percboard{
		name: name,
		sz:   sz,
		con:  make(map[int]bool),
		uf:   union.NewIntUnionFindPCQU(sz*sz + 3), // here we add 2 virtual points that connect top and bottom rows
		uf1:  union.NewIntUnionFindPCQU(sz*sz + 3), // this is to fight backwash which might happen from isFull not marking correctly
		mut:  &sync.Mutex{},
	}
}

func (p *percboard) connectVirtualPoints() { //we don't connect the last line with uf1
	dim := p.sz * p.sz
	top, bot := dim+1, dim+2
	p.uf.Union(top, 1)
	p.uf.Union(bot, dim-1)
	p.uf1.Union(top, 1)
	for i := 0; i < p.sz; i++ {
		p.uf.Union(1, i+1)
		p.uf.Union(dim-1, dim-i)
		p.uf1.Union(1, i+1)
	}
}

func (p *percboard) isOpen(x, y int) bool {
	val := (x-1)*p.sz + y
	p.mut.Lock()
	_, ok := p.con[val]
	defer p.mut.Unlock()
	return ok
}

func (p *percappctx) viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><h2> Boards registered </h2><br> <div style='color: #040'>")
	for k, _ := range p.boards {
		fmt.Fprintf(w, "%s<br>", k)
	}
	fmt.Fprintf(w, "</div></html>")
}

func (p *percappctx) createHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	size, ok1 := vars["size"]
	if !ok || !ok1 {
		fmt.Printf("error: var name or size non existent\n")
		return
	}
	if _, ok := p.boards[name]; ok {
		fmt.Printf("error: name exists\n")
		return
	}
	sz, err := strconv.ParseInt(size, 10, 32)
	if err != nil {
		fmt.Printf("conversion err:%s\n", err)
		return
	}
	p.boards[name] = newPercBoard(name, int(sz))
	p.boards[name].connectVirtualPoints()
}

func drawCanvas(b *percboard) string {
	var (
		ok bool
	)
	s1 := `
<html>
 <head>
  <script type="text/javascript">
    function draw() {
      var canvas = document.getElementById('canvas');
      if (canvas.getContext) {
        var ctx = canvas.getContext('2d')
`
	s2 := `
      }
    }
    </script>
 </head>
 <body onload="draw();">
   <canvas id="canvas" width="500" height="500"></canvas>
`
	s3 := `
 </body>
</html>`

	bstr := ""
	x, y := boxspacing, boxspacing
	top := b.sz*b.sz + 1 //the virtual points
	bot := top + 1
	for i := 0; i < top-1; i++ {
		if i%b.sz == 0 && i != 0 { //increase x zero y if not the first time
			y += boxsz + boxspacing
			x = boxspacing
		}
		contop, errtop := b.uf1.IsConnected(i+1, top) //only check in uf1 to fight backwash
		if errtop != nil {
			return fmt.Sprintf("error in IsConnected :%s\n", errtop)
		}

		if _, ok = b.con[i+1]; ok && !contop {
			bstr += "ctx.fillStyle = 'rgb(200,0,0)';\n" //red has been opened
		} else if ok && contop {
			bstr += "ctx.fillStyle = 'rgb(0,200,0)';\n" //green is full
		} else {
			bstr += "ctx.fillStyle = 'rgb(0,0,0)';\n" //black not open
		}
		bstr += fmt.Sprintf("ctx.fillRect(%d,%d,%d,%d);\n", x, y, boxsz, boxsz)
		x += boxsz + boxspacing
	}
	resstr := ""
	rescon, _ := b.uf.IsConnected(top, bot) //percolation test
	if rescon {
		resstr += "<br>percolates<br>"
	} else {
		resstr += "<br>doesn't percolate<br>"
	}
	return s1 + bstr + s2 + resstr + s3
}

func (p *percappctx) viewNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		fmt.Printf("error: name non existent\n")
		return
	}
	if b, ok := p.boards[name]; !ok {

		fmt.Printf("error: name non existent\n")
		return
	} else {
		b.mut.Lock()
		fmt.Fprintf(w, drawCanvas(b))
		b.mut.Unlock()
	}
	return
}

func (p *percappctx) openHandler(w http.ResponseWriter, r *http.Request) {
	var (
		b  *percboard
		ok bool
	)
	vars := mux.Vars(r)
	name, ok := vars["name"]
	point1, ok1 := vars["point1"]
	point2, ok2 := vars["point2"]
	if !ok || !ok1 || !ok2 {
		fmt.Printf("error: var name or point1 or point2 non existent\n")
		return
	}
	if b, ok = p.boards[name]; !ok {
		fmt.Printf("unknown board name\n")
		return

	}
	irow, err1 := strconv.ParseInt(point1, 10, 32)
	icol, err2 := strconv.ParseInt(point2, 10, 32)
	if err1 != nil || err2 != nil {
		fmt.Printf("convertion error: p1:%s p2:%s\n", err1, err2)
		return
	}
	row := int(irow)
	col := int(icol)
	val := int(row-1)*b.sz + int(col)
	upval := val - b.sz
	downval := val + b.sz
	rightval := val + 1
	leftval := val - 1
	if row > 1 && b.isOpen(row-1, col) {
		b.uf.Union(val, upval)
		b.uf1.Union(val, upval)
	}
	if row < b.sz && b.isOpen(row+1, col) {
		b.uf.Union(val, downval)
		b.uf1.Union(val, downval)
	}
	if col > 1 && b.isOpen(row, col-1) {
		b.uf.Union(val, leftval)
		b.uf1.Union(val, leftval)
	}
	if col < b.sz && b.isOpen(row, col+1) {
		b.uf.Union(val, rightval)
		b.uf1.Union(val, rightval)
	}
	b.mut.Lock()
	b.con[val] = true
	b.mut.Unlock()
	fmt.Printf("open %d %d\n", row, col)
}

func main() {
	ctx := newPercAppCtx()
	ctx.router.HandleFunc("/view/", ctx.viewHandler).Methods("GET")
	ctx.router.HandleFunc("/create/{name}/{size}/", ctx.createHandler).Methods("GET")
	ctx.router.HandleFunc("/open/{name}/{point1}/{point2}/", ctx.openHandler).Methods("GET")
	ctx.router.HandleFunc("/view/{name}/", ctx.viewNameHandler).Methods("GET")
	http.ListenAndServe(":8080", ctx.router)

}
