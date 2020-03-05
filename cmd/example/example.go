package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"

	"github.com/zetamatta/go-windows10-ansi"
)

func main() {
	fmt.Println("Defaut-Mode: \x1B[32;1m!!!!\x1B[0m\n")

	closer, err := ansi.EnableStdoutVirtualTerminalProcessing()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, reflect.TypeOf(err).String())
		if errno, ok := err.(syscall.Errno); ok {
			fmt.Fprintf(os.Stderr, "syscall.Errno(%d)\n", errno)
		}
		os.Exit(1)
	}

	fmt.Println("Ansi-Mode: \x1B[32;1m!!!!\x1B[0m\n")

	closer()

	fmt.Println("Default-Mode: \x1B[32;1m!!!!\x1B[0m\n")
}
