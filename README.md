<h1 align="center">Goggle 🥽</h1>
<p align="center">Search your api through types, with speed </p>

## Overview

Goggle is a type-directed search engine like [hoogle](https://github.com/ndmitchell/hoogle) but for [Go](https://go.dev/)

## :tada: Try Goggle now!

Try Goggle for yourself! You can now visit [here](https://abiriadev.github.io/goggle/) to see Goggle in action.

## :camera: Demo

![](./assets/demo.png)

## :mag: Query

You can type query to search and filter results.

The most simplest form is just Go's standard function definition.

```go
func length(s string) int
```

But we can omit a function name, to retrieve results whose name does not match with `length`.

```go
func (s string) int
```

We can omit a `func` keyword too.

```go
(s string) int
```

Finally, we can omit argument names.

```go
(string) int
```

### Query syntax definition

```ebnf
Type = Primitives | SliceType | PointerType | identifier .
Primitives = "bool" | Int | UInt | "float32" | "float64" | "complex64" | "complex128" .
Int = "int" | "int8" | "int16" | "int32" | "int64" .
UInt = "uint" | "uint8" | "uint16" | "uint32" | "uint64" | "uintptr" .

SliceType = "[" "]" Type .

Parameters = "(" [ Type { "," Type } ] ")" .
Query = [ "func" ] [ identifier ] Parameters [ Type ] .
```

## Build Manually

```sh
$ git clone https://github.com/abiriadev/goggle && cd goggle
```

### Build indexer from source and index custom packages

```sh
$ go run ./cmd/indexer
# or
$ go run ./cmd/indexer <space separated list of packages to index>
```

See help page for more information:

```sh
Usage of indexer:
  -f string
        index format (default "gob")
  -o string
        path to save index file
```

### Build and run REPL

```sh
$ go run ./cmd/repl
# or optionally pass a path to index file to use
$ go run ./cmd/repl <index file to use>
```

It will then show you a prompt starting with `λ`.

Type any query(like `() bool`) and enter to see the results.

```go
λ () bool
func utf8.FullRune() bool       // FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.
func nettest.TestableAddress() bool     // TestableAddress reports whether address of network is testable on the current platform configuration.
func nettest.SupportsRawSocket() bool   // SupportsRawSocket reports whether the current session is available to use raw sockets.
func nettest.SupportsIPv6() bool        // SupportsIPv6 reports whether the platform supports IPv6 networking functionality.
func nettest.SupportsIPv4() bool        // SupportsIPv4 reports whether the platform supports IPv4 networking functionality.
func signal.Ignored() bool      // Ignored reports whether sig is currently ignored.
func slices.Equal() bool        // Equal reports whether two slices are equal: the same length and all elements equal.
func testenv.OptimizationOff() bool     // OptimizationOff reports whether optimization is disabled.
func testenv.HasSymlink() bool  // HasSymlink reports whether the current system can use os.Symlink.
func testenv.HasSrc() bool      // HasSrc reports whether the entire source tree is available under GOROOT.
```

### Build and run Goggle server

```sh
$ go run ./cmd/goggle
```

The default port number is `6099`(L33T or `Gogg`). You can pass `-port` option to change it.

```sh
Usage of goggle:
  -port int
        port number to bind (default 6099)
```

Try requesting from terminal:

```sh
$ http :6099/search q='() bool' -v

POST /search HTTP/1.1
Accept: application/json, */*;q=0.5
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Content-Length: 15
Content-Type: application/json
Host: localhost:6099
User-Agent: HTTPie/3.2.1

{
    "q": "() bool"
}

HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Length: 1970
Content-Type: text/plain; charset=utf-8
Date: Tue, 12 Dec 2023 04:12:01 GMT

{
    "items": [
        {
            "sim": 0,
            "sig": "func utf8.FullRune() bool",
            "summary": "FullRune reports whether the bytes in p begin with a full UTF-8 encoding of a rune.",
            "link": "https://pkg.go.dev/unicode/utf8#FullRune"
        },
        {
            "sim": 0,
            "sig": "func nettest.TestableAddress() bool",
            "summary": "TestableAddress reports whether address of network is testable on the current platform configuration.",
            "link": "https://pkg.go.dev/golang.org/x/net/nettest#TestableAddress"
        },
        ...
    ]
}
```

### Build and run frontend

Ensure that you have [Go](https://go.dev), [Task](https://github.com/go-task/task), [Node.js](https://nodejs.org), and [Binaryen](https://github.com/WebAssembly/binaryen) installed.

Then, execuate the following commands:

```sh
$ task wasm-exec syntaxck
$ corepack enable
$ pnpm install --frozen-lockfile
$ cd frontend
```

If you don't want to have local Goggle proxy, you can specify your already-deployed endpoint by setting `VITE_EXTERN_ENDPOINT` variable.

```sh
$ echo 'VITE_EXTERN_ENDPOINT=<type your endpoint url here>' > .env.production
```

Then, run!

```sh
$ pnpm dev
# Or, to use an external endpoint:
$ pnpm dev --mode production
```

For building the frontend for deployment or serving:

```sh
$ pnpm build
```

## :memo: TODO

-   [ ] Index
    -   [x] Portable index file
    -   [ ] Index popular packages
-   [ ] Incremental search
-   [ ] Frontend
    -   [ ] Standalone result view
    -   [x] Link to pkg.go.dev
    -   [x] Brief description
    -   [ ] Syntax hightlighting for search result
    -   [ ] Use dedicated search bar component
-   [ ] Query
    -   [ ] Compound types
        -   [ ] Array
        -   [ ] Slice
        -   [ ] Pointer type
        -   [ ] Inline struct type
        -   [ ] Interface resolution
    -   [ ] Method
    -   [ ] Multiple return
    -   [ ] Parameter collaping
    -   [ ] Spread syntax
    -   [ ] Generics
        -   [ ] Constraints
-   [x] Levenshtein distance
    -   [ ] Argument-level similarity comparison
    -   [ ] Hoogle-like structured edit distance
        -   [ ] Subtype polymorphic edit distance
-   [x] GHA CD automation
-   [ ] External tools
    -   [x] REPL
    -   [ ] vscode extension
    -   [ ] neovim LSP support?

## :grinning: This is really awwwesome!! How can I help?

There are many ways to support and contribute to the ongoing maintenance and improvement of Goggle. Any support is greatly appreciated!

-   **Spread the world.** Share Goggle with your co-workers, students, and community so that they can find it useful as well!
-   **Report bugs.** If you encounter any unexpected behavior or runtime panics, please open an issue to report and document them.
-   **Make your document cleaner.** Although Goggle can find items without documentation, it doesn't have a power to generate intuitive identifiers and descriptive summaries. So it's a good idea to document you package thoroughly to enhance developer experience.
-   **Suggest better idea.** Currently, Goggle's approximate search doesn't support structural edit-distance, and there are still a lot of missing features. Your suggestions for more accurate and efficient implementations are always welcome.
-   **Build creative tools on top of Goggle.** Goggle currently supports web search and REPL, but the possibilities for its application are limitless. Ideas like a vscode extension, LSP autocompletion, etc., could significantly expand its ecosystem.
