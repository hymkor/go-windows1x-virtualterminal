//go:build run
// +build run

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
