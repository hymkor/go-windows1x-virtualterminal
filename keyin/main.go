package keyin

import (
	"os"

	"golang.org/x/term"
)

// Raw switches the terminal to the raw mode. To restore it the cooked mode,
// the function given with the first return value has to be called.
// It is the wrapper of `term.MakeRow(int(os.Stdin.Fd()))` and
// `term.Restore(int(os.Stdin.Fd()),...)`.
func Raw() (func(), error) {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return func() {}, err
	}
	return func() { term.Restore(int(os.Stdin.Fd()), state) }, nil
}

func Get() (string, error) {
	var buffer [256]byte

	n, err := os.Stdin.Read(buffer[:])
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
