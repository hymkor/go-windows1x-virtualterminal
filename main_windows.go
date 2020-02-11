// +build windows

package ansi

import (
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

const enableVirtualTerminalProcessing uint32 = 0x0004

func enableStdoutVirtualTerminalProcessing() (func(), error) {
	return changeConsoleMode(windows.Stdout, ModeSet(enableVirtualTerminalProcessing))
}

func enableStderrVirtualTerminalProcessing() (func(), error) {
	return changeConsoleMode(windows.Stderr, ModeSet(enableVirtualTerminalProcessing))
}
