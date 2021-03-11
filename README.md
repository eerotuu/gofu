# Go Functional:

### Install

```
go get github.com/eerotuu/gofu/

```

### Quick Start

```go

import "github.com/eerotuu/gofu/arrays


data := arrays.Str{"first", "second", "filter", "last"}

result := data.Filter(func(s string) bool {
  return s != "filter"
}).DropLast().Join()

// result = "firstsecond"

```
