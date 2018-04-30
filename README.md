# wind
A go wrapper of dll for python interface of wind financial terminal.

This program can only be run on Windows. You must have a valid wind and python addin installed.

A short test.

```go
package main

import (
	"fmt"

	"naturebridge-asset.com/wind"
)

func main() {
	wind.Start()
	fmt.Println(wind.Connected())
}
```
