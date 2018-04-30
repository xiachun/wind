# wind
A go wrapper of python addin dll of wind financial terminal.

This package can only be used on Windows. You must have a valid wind and python addin installed.

A short test.

```go
package main

import (
	"fmt"

	"github.com/xiachun/wind"
)

func main() {
	wind.Start()
	fmt.Println(wind.Connected())
}
```
