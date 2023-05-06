//go:build !windows
// +build !windows

package ansi

func enableVirtualTerminalInput() (func(), error) {
	return func() {}, &ErrNotWindows{}
}

func enableStdoutVirtualTerminalProcessing() (func(), error) {
	return func() {}, &ErrNotWindows{}
}

func enableStderrVirtualTerminalProcessing() (func(), error) {
	return func() {}, &ErrNotWindows{}
}
