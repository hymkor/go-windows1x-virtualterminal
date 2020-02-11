package ansi

// EnableStdoutVirtualTerminalProcessing enables Windows10's native ESCAPE SEQUENCE support on STDOUT
func EnableStdoutVirtualTerminalProcessing() (func(), error) {
	return enableStdoutVirtualTerminalProcessing()
}

// EnableStderrVirtualTerminalProcessing enables Windows10's native ESCAPE SEQUENCE support on STDERR
func EnableStderrVirtualTerminalProcessing() (func(), error) {
	return enableStderrVirtualTerminalProcessing()
}
