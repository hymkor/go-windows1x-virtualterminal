package main

import (
	"fmt"
	"os"

	"github.com/zetamatta/go-windows10-ansi"
)

func main() {
	fmt.Println("Defaut-Mode: \x1B[32;1m!!!!\x1B[0m\n")

	closer, err := ansi.EnableStdout()
	if err != nil {
		if ansi.IsNotWindows(err) {
			fmt.Fprintln(os.Stderr, "This machine is not Windows.")
		} else if ansi.IsNotSupported(err) {
			fmt.Fprintln(os.Stderr, "This machine is Windows 8.1 or Windows Server")
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		os.Exit(1)
	}

	fmt.Println("Ansi-Mode: \x1B[32;1m!!!!\x1B[0m\n")

	closer()

	fmt.Println("Default-Mode: \x1B[32;1m!!!!\x1B[0m\n")
}
