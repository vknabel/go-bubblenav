package nav

import tea "github.com/charmbracelet/bubbletea"

type NavPage struct {
	windowSize *tea.WindowSizeMsg
	previous   tea.Model
	current    tea.Model

	disabledQuitKeyBindings bool
}

func NewPage(current tea.Model) NavPage {
	return NavPage{
		windowSize: nil,
		previous:   nil,
		current:    current,
	}
}

func (m NavPage) Init() tea.Cmd {
	if m.current == nil {
		return nil
	}
	return tea.Batch(m.current.Init(), func() tea.Msg {
		if m.windowSize != nil {
			return *m.windowSize
		}
		return nil
	})
}

func (m NavPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = &msg
	case PagePopMsg:
		return m.Pop()
	case PagePushMsg:
		return m.Push(msg.Page)
	case PageReplaceMsg:
		new := NewPage(msg.Page)
		new.previous = m.previous
		new.windowSize = m.windowSize
		new.disabledQuitKeyBindings = m.disabledQuitKeyBindings
		return new, new.Init()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m.Pop()
		case "q":
			if !m.disabledQuitKeyBindings {
				return m, tea.Quit
			}
		}
	}
	m.current, cmd = m.current.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m NavPage) View() string {
	if m.current == nil {
		return ""
	}
	return m.current.View()
}

func (m *NavPage) DisableQuitKeyBindings() {
	m.disabledQuitKeyBindings = true
}

func (m NavPage) Top() tea.Model {
	return m.current
}

func (m NavPage) Pop() (tea.Model, tea.Cmd) {
	if !m.disabledQuitKeyBindings && m.previous == nil {
		return m, tea.Quit
	}
	if m.previous == nil {
		return m, nil
	}
	next := m.previous
	if m.windowSize != nil {
		cmds := []tea.Cmd{}
		var cmd tea.Cmd
		next, cmd = next.Update(*m.windowSize)
		cmds = append(cmds, cmd)
		next, cmd = next.Update(PageRestoreMsg{})
		cmds = append(cmds, cmd)
		return next, tea.Batch(cmds...)
	}
	return next, nil
}

func (m NavPage) Push(page tea.Model) (NavPage, tea.Cmd) {
	new := NewPage(page)
	new.previous = m
	new.windowSize = m.windowSize
	new.disabledQuitKeyBindings = m.disabledQuitKeyBindings
	return new, new.Init()
}
