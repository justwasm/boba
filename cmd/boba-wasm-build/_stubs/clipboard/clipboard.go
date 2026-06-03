// Package clipboard read/write on clipboard
package clipboard

import "os"

var isSsh = os.Getenv("SSH_TTY") != ""

// ReadAll read string from clipboard
func ReadAll() (string, error) {
	return readAll()
}

// WriteAll write string to clipboard
func WriteAll(text string) error {
	if isSsh {
		return writeOsc52(text)
	}
	return writeAll(text)
}
