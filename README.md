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
