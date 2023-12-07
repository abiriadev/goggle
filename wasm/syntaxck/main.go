//go:build wasm

package main

import (
	"syscall/js"

	"github.com/abiriadev/goggle/pkg/query"
)

// global-lifetime query parser
// TODO: use closure later?
var qp = query.QueryParserUnsafe()

func main() {
	js.Global().Set("syntaxck", js.FuncOf(syntaxckSys))

	<-make(chan struct{}, 0)
}

// string -> bool
func syntaxckSys(_ js.Value, args []js.Value) any {
	// SAFETY: check array boundaries
	if len(args) != 1 {
		return "only one argument needed"
	}

	return syntaxck(args[0].String())
}

func syntaxck(query string) bool {
	_, err := qp.ParseString("", query)
	return err == nil
}
