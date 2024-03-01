//go:build run

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hymkor/go-windows1x-virtualterminal"
	"github.com/hymkor/go-windows1x-virtualterminal/keyin"
)

func mains() error {
	// EnableStdin should be called once in one application.
	close1, err := virtualterminal.EnableStdin()
	if err != nil {
		return fmt.Errorf("EnableStdin: %w", err)
	}
	defer close1()

	// Raw switches the terminal to raw-mode.
	// It should be called once in one readline loop.
	close2, err := keyin.Raw()
	if err != nil {
		return fmt.Errorf("Raw: %w", err)
	}
	fmt.Print("Hit any key: ")

	ch, err := keyin.Get()
	if err != nil {
		return err
	}
	close2()

	ch = strings.ReplaceAll(ch, "\x1B", "<ESC>")
	fmt.Println(ch)
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
