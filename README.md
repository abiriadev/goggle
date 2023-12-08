<h1 align="center">Goggle 🥽</h1>
<p align="center">Search your api through types, with speed </p>

Type-directed search engine like [hoogle](https://github.com/ndmitchell/hoogle) but for [Go](https://go.dev/)

## Query syntax overview

```ebnf
Type = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName = identifier | QualifiedIdent .
TypeArgs = "[" TypeList [ ", "] "]" .
TypeList = Type { "," Type } .
TypeLit = ArrayType | StructType | PointerType | FunctionType | InterfaceType | SliceType | MapType | ChannelType

ArrayType = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .

SliceType = "[" "]" ElementType .

FunctionType = "func" Signature .
Signature = Parameters [ Result ]
Result = Parameters | Type .
Parameters = "(" [ ParameterList [ "," ] ] ")" .
ParameterList = ParameterDecl { "," ParameterDecl } .
ParameterDecl = [ IdentifierList ] [ "... " ] Type .

Query = [ "func" ] [ identifier ] Signature .
```

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

## TODO

-   [ ] Generics
    -   [ ] Constraints
-   [x] Portable index file
-   [ ] Incremental search
-   [x] Link to pkg.go.dev
-   [ ] Brief description
-   [ ] Query
    -   [ ] Method
    -   [ ] Multiple return
    -   [ ] Parameter collaping
    -   [ ] Spread syntax
-   [ ] Levenshtein distance
-   [ ] Syntax hightlighting for search result

\*? \_? ?? any? ? ...?

## This is really awwwesome!! How can I help?

There are many ways to support and contribute to the ongoing maintenance and improvement of Goggle. Any support is greatly appreciated!

-   **Spread the world.** Share Goggle with your co-workers, students, and community so that they can find it useful as well!
-   **Report bugs.** If you encounter any unexpected behavior or runtime panics, please open an issue to report and document them.
-   **Make your document cleaner.** Although Goggle can find items without documentation, it doesn't have a power to generate intuitive identifiers and descriptive summaries. So it's a good idea to document you package thoroughly to enhance developer experience.
-   **Suggest better idea.** Currently, Goggle's approximate search doesn't support structural edit-distance, and there are still a lot of missing features. Your suggestions for more accurate and efficient implementations are always welcome.
-   **Build creative tools on top of Goggle.** Goggle currently supports web search and REPL, but the possibilities for its application are limitless. Ideas like a vscode extension, LSP autocompletion, etc., could significantly expand its ecosystem.
