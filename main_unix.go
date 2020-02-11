// +build !windows

package ansi

func enableStdoutVirtualTerminalProcessing() (func(), error) {
	return func() {}, nil
}

func enableStderrVirtualTerminalProcessing() (func(), error) {
	return func() {}, nil
}
