//go:build !windows

package keyin

import (
	"os"

	"github.com/hymkor/go-lazy"
)

var ConInHandle = lazy.Two[uintptr, error]{
	New: func() (uintptr, error) {
		return 1, nil
	},
}

var ConIn = lazy.Two[*os.File, error]{
	New: func() (*os.File, error) {
		return os.Open("/dev/tty")
	},
}
