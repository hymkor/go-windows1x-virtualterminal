//go:build !windows
// +build !windows

package virtualterminal

func enableStdin() (func(), error) {
	return func() {}, nil
}

func enableStdout() (func(), error) {
	return func() {}, nil
}

func enableStderr() (func(), error) {
	return func() {}, nil
}
