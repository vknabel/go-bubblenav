package nav

import tea "github.com/charmbracelet/bubbletea"

type PageRestoreMsg struct{}
type PagePopMsg struct{}
type PagePushMsg struct {
	Page tea.Model
}
type PageReplaceMsg struct {
	Page tea.Model
}

func Pop() tea.Cmd {
	return func() tea.Msg {
		return PagePopMsg{}
	}
}

func Push(page tea.Model) tea.Cmd {
	return func() tea.Msg {
		return PagePushMsg{page}
	}
}

func Replace(page tea.Model) tea.Cmd {
	return func() tea.Msg {
		return PageReplaceMsg{page}
	}
}
