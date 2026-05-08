package boba

import (
	tea "charm.land/bubbletea/v2"
)

// Program provides an abstraction over wasm.Program and tea.Program
type Program interface {
    Kill()
    Printf(template string, args ...any)
    Println(args ...any)
    Quit()
    ReleaseTerminal() error
    RestoreTerminal() error
    Run() (tea.Model, error)
    Send(msg tea.Msg)
    Wait()
}
