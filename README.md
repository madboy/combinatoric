## combinatoric

Combinatoric tools based on python itertools

### example

```go
package main

import (
    "github.com/madboy/combinatoric"
    "fmt"
)

func main() {
    l := []int{1, 3, 5}

    fmt.Println(combinatoric.Combinations(l, 2))
    fmt.Println(combinatoric.Permutations(l, 3))
    fmt.Println(combinatoric.Permutations(l, 2))
}
```

```
[[1 3] [1 5] [3 5]]
[[1 3 5] [3 1 5] [5 1 3] [1 5 3] [3 5 1] [5 3 1]]
[[1 3] [1 5] [3 1] [3 5] [5 1] [5 3]
```
