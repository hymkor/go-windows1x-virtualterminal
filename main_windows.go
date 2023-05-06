//go:build windows
// +build windows

package ansi

import (
	"syscall"

	"golang.org/x/sys/windows"
)

type ModeOp interface {
	Op(mode uint32) uint32
}

type ModeReset uint32

func (this ModeReset) Op(mode uint32) uint32 {
	return mode &^ uint32(this)
}

type ModeSet uint32

func (this ModeSet) Op(mode uint32) uint32 {
	return mode | uint32(this)
}

func changeConsoleMode(console windows.Handle, ops ...ModeOp) (func(), error) {
	var mode uint32
	err := windows.GetConsoleMode(console, &mode)
	if err != nil {
		return func() {}, err
	}
	restore := func() { windows.SetConsoleMode(console, mode) }

	if len(ops) > 0 {
		newMode := mode
		for _, op1 := range ops {
			newMode = op1.Op(newMode)
		}
		err = windows.SetConsoleMode(console, newMode)
	}
	return restore, err
}

const _PARAMETER_IS_INCORRECT = 87

func enableVirtualTerminalInput() (func(), error) {
	f, err := changeConsoleMode(windows.Stdin, ModeSet(windows.ENABLE_VIRTUAL_TERMINAL_INPUT))
	if errno, ok := err.(syscall.Errno); ok && errno == _PARAMETER_IS_INCORRECT {
		return f, ErrNotSupported
	}
	return f, err
}

func enableVirtualTerminalProcessing(h windows.Handle) (func(), error) {
	f, err := changeConsoleMode(h, ModeSet(windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING))
	if errno, ok := err.(syscall.Errno); ok && errno == _PARAMETER_IS_INCORRECT {
		return f, ErrNotSupported
	}
	return f, err
}

func enableStdoutVirtualTerminalProcessing() (func(), error) {
	return enableVirtualTerminalProcessing(windows.Stdout)
}

func enableStderrVirtualTerminalProcessing() (func(), error) {
	return enableVirtualTerminalProcessing(windows.Stderr)
}
