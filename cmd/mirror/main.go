package main

import (
	"encoding/json"
	"os"

	"github.com/abiriadev/goggle/pkg/index"
)

func main() {
	m, err := index.ResolveFullIndex()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("mirror.json")
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(f).Encode(m)
	if err != nil {
		panic(err)
	}
}
