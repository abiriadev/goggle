package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  byte
	pin  int
}

type res1 struct {
	Page    int `json:"page"`
	Fruites []string
}

func main() {
	p1 := Person{
		Name: "Abiria",
		Age:  17,
		pin:  1234,
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		panic(err)
	}

	fmt.Println(bs)

	var p2 Person

	e2 := json.Unmarshal(bs, &p2)
	if e2 != nil {
		panic(e2)
	}

	fmt.Println(p2)

	a, _ := json.Marshal(true)
	fmt.Println(string(a))

	b, _ := json.Marshal([]string{"123", "abc"})
	fmt.Println(string(b))

	c, _ := json.Marshal(map[string]int{"apple": 5, "abc": 123})
	fmt.Println(string(c))

	d, _ := json.Marshal(res1{
		Page:    123,
		Fruites: []string{"apple", "peach"},
	})
	fmt.Println(string(d))
}
