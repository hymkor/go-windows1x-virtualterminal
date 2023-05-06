//go:build windows
// +build windows

package ansi

import (
	"syscall"

	"golang.org/x/sys/windows"
)

const _PARAMETER_IS_INCORRECT = 87

func changeConsoleMode(console windows.Handle, bits uint32) (func(), error) {
	var mode uint32
	err := windows.GetConsoleMode(console, &mode)
	if err != nil {
		return func() {}, err
	}
	err = windows.SetConsoleMode(console, mode|bits)
	if errno, ok := err.(syscall.Errno); ok && errno == _PARAMETER_IS_INCORRECT {
		err = ErrNotSupported
	}
	return func() { windows.SetConsoleMode(console, mode) }, err
}

func enableStdin() (func(), error) {
	return changeConsoleMode(windows.Stdin, windows.ENABLE_VIRTUAL_TERMINAL_INPUT)
}

func enableStdout() (func(), error) {
	return changeConsoleMode(windows.Stdout, windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

func enableStderr() (func(), error) {
	return changeConsoleMode(windows.Stderr, windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
