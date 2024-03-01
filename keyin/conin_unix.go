//go:build !windows

package keyin

import (
	"os"
	"sync"
)

func ConInHandle() (uintptr, error) {
	return 1, nil
}

var ConIn = sync.OnceValues(func() (*os.File, error) {
	return os.Open("/dev/tty")
})
