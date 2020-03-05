package ansi

type ErrNotSupportVirtualTerminalProcessing struct {
	err error
}

func (e ErrNotSupportVirtualTerminalProcessing) Error() string {
	return "Not support Virtual-Terminal-Processing"
}

func (e ErrNotSupportVirtualTerminalProcessing) Unwrap() error {
	return e.err
}

type ErrNotWindows struct{}

func (e ErrNotWindows) Error() string {
	return "Not Windows"
}

// EnableStdoutVirtualTerminalProcessing enables Windows10's native ESCAPE SEQUENCE support on STDOUT
func EnableStdoutVirtualTerminalProcessing() (func(), error) {
	return enableStdoutVirtualTerminalProcessing()
}

// EnableStderrVirtualTerminalProcessing enables Windows10's native ESCAPE SEQUENCE support on STDERR
func EnableStderrVirtualTerminalProcessing() (func(), error) {
	return enableStderrVirtualTerminalProcessing()
}
