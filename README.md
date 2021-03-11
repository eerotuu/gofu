# Go Functional:

### Quick Start

```go

import "github.com/eerotuu/gofu/arrays


data = Str{"first", "second", "filter", "last"}

result := data.Filter(func(s string) bool {
	return s != "filter"
}).DropLast().Join()

// result = "firstsecond"

```
