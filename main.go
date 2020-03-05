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

func IsNotSupported(err error) bool {
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

// EnableStdout enables Windows10's native ESCAPE SEQUENCE support on STDOUT
func EnableStdout() (func(), error) {
	return enableStdoutVirtualTerminalProcessing()
}

// EnableStdoutVirtualTerminalProcessing is same with EnableStdout for compatible.
func EnableStdoutVirtualTerminalProcessing() (func(), error) {
	return enableStdoutVirtualTerminalProcessing()
}

// EnableStderr enables Windows10's native ESCAPE SEQUENCE support on STDERR
func EnableStderr() (func(), error) {
	return enableStderrVirtualTerminalProcessing()
}

// EnableStderrVirtualTerminalProcessing is same with EnableStderr for compatible
func EnableStderrVirtualTerminalProcessing() (func(), error) {
	return enableStderrVirtualTerminalProcessing()
}
