package main

import (
	"flag"
	"fmt"

	"github.com/abiriadev/goggle/pkg/goggle"
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

	goggle, err := goggle.Load(f)
	if err != nil {
		panic(err)
	}

	rl, err := readline.New("λ ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		rs, err := goggle.Query(line)
		if err != nil {
			fmt.Println(err)
		}

		for _, ri := range rs.Results {
			fmt.Println(ri)
		}
	}
}
