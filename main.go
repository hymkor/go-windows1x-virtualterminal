package ansi

type ErrNotSupportVirtualTerminalInput struct {
	err error
}

func (e ErrNotSupportVirtualTerminalInput) Error() string {
	return "Not support Virtual-Terminal-Input"
}

func (e ErrNotSupportVirtualTerminalInput) Unwrap() error {
	return e.err
}

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
	if _, ok := err.(*ErrNotSupportVirtualTerminalInput); ok {
		return true
	}
	if _, ok := err.(*ErrNotSupportVirtualTerminalProcessing); ok {
		return true
	}
	return false
}

type ErrNotWindows struct{}

func (e ErrNotWindows) Error() string {
	return "Not Windows"
}

func IsNotWindows(err error) bool {
	_, ok := err.(*ErrNotWindows)
	return ok
}

func EnableStdin() (func(), error) {
	return enableVirtualTerminalInput()
}

func EnableVirtualTerminalInput() (func(), error) {
	return enableVirtualTerminalInput()
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
