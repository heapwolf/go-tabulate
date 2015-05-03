# SYNOPSIS
Column-like layout (padded to fit) for arrays of strings.

Each string from a `[]string` is padded to fit a given width
and a string will be returned. Useful with [`termsize`][0].

# EXAMPLE
```go
package main

import "fmt"
import "github.com/hij1nx/go-tabulate"

var w = 80 // usually the calculated terminal width
var a = []string{"foo", "barrrrrr", "bazz", "blaaaaa"}

fmt.Println(Tabulate(w, a))
```

# SCREENSHOT
![img](/cols.png)

[0]:https://github.com/hij1nx/go-termsize

