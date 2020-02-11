go-windows10-ansi
=================

This package enables Windows10's escape sequences.

```
// +build example

package main

import (
    "fmt"
    "os"

    "github.com/zetamatta/go-windows10-ansi"
)

func main() {
    fmt.Println("Defaut-Mode: \x1B[32;1m!!!!\x1B[0m\n")

    closer, err := ansi.EnableStdoutVirtualTerminalProcessing()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }

    fmt.Println("Ansi-Mode: \x1B[32;1m!!!!\x1B[0m\n")

    closer()

    fmt.Println("Default-Mode: \x1B[32;1m!!!!\x1B[0m\n")
}
```

<img src="example.png" />
