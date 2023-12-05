package main

import (
	"flag"
	"fmt"

	"github.com/abiriadev/goggle/pkg/goggle"
	"github.com/abiriadev/goggle/pkg/query"
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

	index, e := goggle.Parse(f)

	if e != nil {
		panic(e)
	}

	qs := "() bool"

	qp, e := query.QueryParser()
	if e != nil {
		panic(e)
	}

	q, e := qp.ParseString("", qs)
	if e != nil {
		panic(e)
	}

	for _, fd := range index.Query(q) {
		fmt.Println(fd.Name, "->", fd.Ret)
	}
}
