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
