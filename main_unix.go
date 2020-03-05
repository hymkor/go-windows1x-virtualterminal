// +build !windows

package ansi

func enableStdoutVirtualTerminalProcessing() (func(), error) {
	return func() {}, &ErrNotWindows{}
}

func enableStderrVirtualTerminalProcessing() (func(), error) {
	return func() {}, &ErrNotWindows{}
}
