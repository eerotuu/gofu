# Go Functional:

Created this module for learning Golang and its possibilities as functional language.
**This probably is something you shouldn't use in real project.**

### Install

```
go get github.com/eerotuu/gofu/
```

### Quick Start

```go
import "github.com/eerotuu/gofu/arrays


data := arrays.Str{"hello", "not wanted", "world!", "last"}

result := data.Filter(func(s string) bool {
  return s != "not wanted"
}).DropLast().Map(func(s string) string {
  return strings.Title(s)
}).Join(" ")

// result = "Hello World!"
```

### Arrays

---

**Types**

- **Str** `[]string`

- **Int** `[]Int`

- **Float** `[]float32`

- **Double** `[]float64`

All types can be converted with each other using type methods.

```go
strArr := arrays.Str{"1", "5", "7"}
intArr := strArr.Int()
floatArr := intArr.Float()
doubleArr := strArr.Double()

// etc...
```

**Generic Functions** `[]string` `[]int` `[]float32` `[]float64`

- Map()
- Filter()
- Reduce()
- DropLast()

**Type Specific Functions:**

**Str** `[]string`

- Join()
