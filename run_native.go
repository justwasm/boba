//go:build !js

package boba

import (
	tea "charm.land/bubbletea/v2"
)

// NewProgram creates a new BubbleTea program for the native terminal.
var NewProgram = tea.NewProgram

// Run executes the given BubbleTea model with the appropriate runtime
// for the build target.
func Run(model tea.Model, opts ...tea.ProgramOption) error {
	_, err := NewProgram(model, opts...).Run()
	return err
}
