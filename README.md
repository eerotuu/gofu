# Go Functional:

### Quick Start

```go

import "github.com/eerotuu/gofu/arrays
data = Str{"first", "remove", "last"}
result := res.Filter(func(s string) bool {
	return s != "remove"
}).DropLast().Join()

```
