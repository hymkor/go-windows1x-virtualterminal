package ansi

import "errors"

var ErrNotSupported = errors.New("Not support Virtual-Terminal")

// EnableStdin enables Windows10's `ENABLE_VIRTUAL_TERMINAL_INPUT` on STDIN
func EnableStdin() (func(), error) { return enableVirtualTerminalInput() }

// EnableStdout enables Windows10's native ESCAPE SEQUENCE support on STDOUT
func EnableStdout() (func(), error) { return enableStdoutVirtualTerminalProcessing() }

// EnableStderr enables Windows10's native ESCAPE SEQUENCE support on STDERR
func EnableStderr() (func(), error) { return enableStderrVirtualTerminalProcessing() }
