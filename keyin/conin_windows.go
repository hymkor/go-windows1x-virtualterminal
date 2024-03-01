package keyin

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"

	"github.com/hymkor/go-lazy"
)

var ConInHandle = lazy.Two[uintptr, error]{
	New: func() (uintptr, error) {
		conin := []uint16{'C', 'O', 'N', 'I', 'N', '$', 0}
		h, err := windows.CreateFile(
			&conin[0],
			windows.GENERIC_READ|windows.GENERIC_WRITE,
			uint32(windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE),
			nil,
			windows.OPEN_EXISTING,
			windows.FILE_ATTRIBUTE_NORMAL,
			0)
		if err != nil {
			return 0, fmt.Errorf("windows.CreateFile: %w", err)
		}
		return uintptr(h), err
	},
}

var ConIn = lazy.Two[*os.File, error]{
	New: func() (*os.File, error) {
		in, err := ConInHandle.Values()
		if err != nil {
			return nil, err
		}
		return os.NewFile(in, "/dev/tty"), nil
	},
}
