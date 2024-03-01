//go:build !windows

package conin

import (
	"os"

	"github.com/hymkor/go-lazy"
)

var Handle = lazy.Two[uintptr, error]{
	New: func() (uintptr, error) {
		return 1, nil
	},
}

var File = lazy.Two[*os.File, error]{
	New: func() (*os.File, error) {
		return os.Open("/dev/tty")
	},
}
