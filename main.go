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

func IsVirtualTerminalProcessingNotSupported(err error) bool {
	_, ok := err.(*ErrNotSupportVirtualTerminalProcessing)
	return ok
}

type ErrNotWindows struct{}

func (e ErrNotWindows) Error() string {
	return "Not Windows"
}

func IsNotWindows(err error) bool {
	_, ok := err.(*ErrNotWindows)
	return ok
}

// EnableStdoutVirtualTerminalProcessing enables Windows10's native ESCAPE SEQUENCE support on STDOUT
func EnableStdoutVirtualTerminalProcessing() (func(), error) {
	return enableStdoutVirtualTerminalProcessing()
}

// EnableStderrVirtualTerminalProcessing enables Windows10's native ESCAPE SEQUENCE support on STDERR
func EnableStderrVirtualTerminalProcessing() (func(), error) {
	return enableStderrVirtualTerminalProcessing()
}
