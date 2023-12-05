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
-   [ ] Portable index file
-   [ ] Incremental search
-   [ ] Link to pkg.go.dev
-   [ ] Brief description
-   [ ] Query
    -   [ ] Method
    -   [ ] Multiple return
    -   [ ] Parameter collaping
    -   [ ] Spread syntax
-   [ ] Levenshtein distance
