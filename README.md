go-windows1x-virtualterminal
=============================

This package enables Windows10 and 11's

- Escape sequences for STDOUT and STDERR
- Arrow-keys and others on `func (*Terminal) ReadLine` of [golang.org/x/term](https://pkg.go.dev/golang.org/x/term)

Enable escape sequences for STDOUT and STDERR
---------------------------------------

```examples/example.go
package main

import (
    "errors"
    "fmt"
    "os"

    "github.com/hymkor/go-windows1x-virtualterminal"
)

func mains() error {
    fmt.Println("Defaut-Mode: \x1B[32;1m!!!!\x1B[0m\n")
    closer, err := virtualterminal.EnableStdout()
    if err != nil {
        if errors.Is(err, virtualterminal.ErrNotSupported) {
            return fmt.Errorf("This machine can not use ANSI-ESCAPE-SEQUENCE: %w", err)
        }
        return err
    }
    fmt.Println("Ansi-Mode: \x1B[32;1m!!!!\x1B[0m\n")
    closer()
    fmt.Println("Default-Mode: \x1B[32;1m!!!!\x1B[0m\n")
    return nil
}

func main() {
    if err := mains(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
```

![example.png](./example.png)

Enable arrow keys and others at `func (*Terminal) ReadLine` of golang.org/x/term
-------------------------------------------------------------------------

```examples/example2.go
package main

import (
    "fmt"
    "io"
    "os"

    "golang.org/x/term"

    "github.com/hymkor/go-windows1x-virtualterminal"
)

func mains() error {
    oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        return err
    }
    defer term.Restore(int(os.Stdin.Fd()), oldState)

    disable, err := virtualterminal.EnableStdin()
    if err != nil {
        return err
    }
    defer disable()

    terminal := term.NewTerminal(&struct {
        io.Reader
        io.Writer
    }{os.Stdin, os.Stdout}, "> ")

    line, err := terminal.ReadLine()
    if err != nil {
        return err
    }
    fmt.Println("Line:", line)
    return nil
}

func main() {
    if err := mains(); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }

}
```
