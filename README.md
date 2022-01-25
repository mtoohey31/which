# which

A Go library implementation of `which`, that tries to follow the same algorithm as the original. Please notify me if you notice any discrepancies between this library's behaviour, and the original's.

## Usage

```go
package main

import "mtoohey.com/which"

func main() {
  println(which.Which("ls"))
}
```
