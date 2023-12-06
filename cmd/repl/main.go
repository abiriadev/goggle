package main

import (
	"flag"
	"fmt"

	"github.com/abiriadev/goggle/pkg/eval"
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
	"github.com/chzyer/readline"
)

func pathArg() string {
	flag.Parse()
	f := flag.Arg(0)

	if f == "" {
		return "index.json"
	}

	return f
}

func main() {
	f := pathArg()

	index, e := index.Load(f)
	if e != nil {
		panic(e)
	}

	rl, e := readline.New("λ ")
	if e != nil {
		panic(e)
	}
	defer rl.Close()

	for {
		line, e := rl.Readline()
		if e != nil {
			break
		}

		qs := line

		qp, e := query.QueryParser()
		if e != nil {
			panic(e)
		}

		q, e := qp.ParseString("", qs)
		if e != nil {
			fmt.Println(e)
		}

		for _, fd := range eval.Query(*q) {
			fmt.Println(fd.String())
		}
	}
}
