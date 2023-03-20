# go-bubblenav

Early development navigation helpers for charm bubbletea.

```go
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vknabel/go-puffery/nav"
)

type appOptions struct {
}

type appModel struct {
	stack   tea.Model
	options AppOptions

	err error
}

func InitialAppModel(options AppOptions) appModel {
	initial := myInitialPageModel()
	return appModel{
		stack: nav.NewPage(initial),
	}
}

func (m appModel) Init() tea.Cmd {
	return m.stack.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.stack, cmd = m.stack.Update(msg) // this will call the update function of the current page
	return m, cmd
}

func (m appModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	return m.stack.View()
}

```
