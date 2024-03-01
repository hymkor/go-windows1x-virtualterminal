//go:build run

package main

import (
	"fmt"
	"os"

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

	for _, c := range ch {
		if c < 0x20 {
			fmt.Printf("<C-%c>", '@'+c)
		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
