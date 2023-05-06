//go:build !windows
// +build !windows

package ansi

func enableVirtualTerminalInput() (func(), error) {
	return func() {}, nil
}

func enableStdoutVirtualTerminalProcessing() (func(), error) {
	return func() {}, nil
}

func enableStderrVirtualTerminalProcessing() (func(), error) {
	return func() {}, nil
}
