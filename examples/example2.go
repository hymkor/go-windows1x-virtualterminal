package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"

	"github.com/hymkor/go-windows1x-virtualterminal"
)

func readLine() (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	disableIn, err := virtualterminal.EnableStdin()
	if err != nil {
		return "", err
	}
	defer disableIn()

	disableOut, err := virtualterminal.EnableStdout()
	if err != nil {
		return "", err
	}
	defer disableOut()

	terminal := term.NewTerminal(&struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}, "> ")

	return terminal.ReadLine()
}

func main() {
	line, err := readLine()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Line:", line)
}
