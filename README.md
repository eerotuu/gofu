# Go Functional:

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

All types can beconverted with each other using type methods.

```go
strArr := arrays.Str{"1", "5", "7"}
intArr := strArr.Int()
floatArr := intArr.Float()
doubleArr := strArr.Double()

// etc...
```

**Generic Functions** `[]string, []int, []float32, []float64`

- Map()
- Filter()
- Reduce()
- DropLast()

**Type Secific Functions:**

**Str** `[]string`

- Join()
