<h1 align="center">Goggle 🥽</h1>
<p align="center">Search your api through types, with speed </p>

Type-directed search engine like [hoogle](https://github.com/ndmitchell/hoogle) but for [Go](https://go.dev/)

## Query syntax overview

```go
// normal function definition
func length(s string) int

// omit function name
func (s string) int

// omit `func` keyword
(s string) int

// multiline parameters
(s string, flag int) int

// multiple return values
(s string) (int, error)

// wildcard
(_) bool

// generics
[T](s T) int
```
