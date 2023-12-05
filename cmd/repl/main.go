package main

import (
	"flag"
	"fmt"

	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
	"github.com/chzyer/readline"
)

func pathArg() string {
	flag.Parse()
	f := flag.Arg(0)

	if f == "" {
		panic("provide path to target package")
	}

	return f
}

func main() {
	f := pathArg()

	index, e := index.IndexPackage(f)

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

		// fmt.Println(line)

		qs := line

		qp, e := query.QueryParser()
		if e != nil {
			panic(e)
		}

		q, e := qp.ParseString("", qs)
		if e != nil {
			fmt.Println(e)
		}

		for _, fd := range index.Query(*q) {
			fmt.Println(fd.Name, "->", fd.Ret)
		}
	}
}
