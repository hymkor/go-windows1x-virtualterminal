package keyin

import (
	"fmt"
	"os"

	"golang.org/x/term"

	"github.com/hymkor/go-windows1x-virtualterminal/conin"
)

// Raw switches the terminal to the raw mode. To restore it the cooked mode,
// the function given with the first return value has to be called.
// It is the wrapper of `term.MakeRow(int(os.Stdin.Fd()))` and
// `term.Restore(int(os.Stdin.Fd()),...)`.
func Raw() (func(), error) {
	conIn, err := conin.Handle.Values()
	if err != nil {
		return func() {}, fmt.Errorf("ConInHandle: %w", err)
	}
	state, err := term.MakeRaw(int(conIn))
	if err != nil {
		return func() {}, fmt.Errorf("term.MakeRaw: %w", err)
	}
	return func() { term.Restore(int(conIn), state) }, nil
}

func Get() (string, error) {
	var buffer [256]byte

	n, err := os.Stdin.Read(buffer[:])
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
