package nav

import tea "github.com/charmbracelet/bubbletea"

type NavStack struct {
	stack      []tea.Model
	windowSize *tea.WindowSizeMsg

	disabledQuitKeyBindings bool
}

func NewStack(root tea.Model) NavStack {
	return NavStack{
		stack:      []tea.Model{root},
		windowSize: nil,
	}
}

func (m *NavStack) DisableQuitKeyBindings() {
	m.disabledQuitKeyBindings = true
}

func (m NavStack) Top() tea.Model {
	if len(m.stack) == 0 {
		return nil
	}
	return m.stack[len(m.stack)-1]
}

func (m *NavStack) Push(page tea.Model) tea.Cmd {
	m.stack = append(m.stack, page)
	return m.initTop()
}

func (m *NavStack) Pop() tea.Cmd {
	if len(m.stack) == 1 {
		if m.disabledQuitKeyBindings {
			return nil
		} else {
			return tea.Quit
		}
	}
	m.stack = m.stack[:len(m.stack)-1]
	if m.windowSize == nil {
		return nil
	}
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	m.stack[len(m.stack)-1], cmd = m.Top().Update(*m.windowSize)
	cmds = append(cmds, cmd)
	m.stack[len(m.stack)-1], cmd = m.Top().Update(PageRestoreMsg{})
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

func (m *NavStack) Replace(page tea.Model) tea.Cmd {
	m.stack[len(m.stack)-1] = page
	return m.initTop()
}

func (m NavStack) Init() tea.Cmd {
	return m.initTop()
}

func (m NavStack) initTop() tea.Cmd {
	cmds := []tea.Cmd{m.Top().Init()}
	if m.windowSize != nil {
		top, cmd := m.Top().Update(*m.windowSize)
		m.stack[len(m.stack)-1] = top
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

func (m NavStack) UpdateModel(msg tea.Msg) (NavStack, tea.Cmd) {
	switch msg := msg.(type) {
	case PagePopMsg:
		return m, m.Pop()
	case PagePushMsg:
		return m, m.Push(msg.Page)
	case PageReplaceMsg:
		return m, m.Replace(msg.Page)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, m.Pop()
		case "q":
			if !m.disabledQuitKeyBindings && len(m.stack) < 2 {
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		m.windowSize = &msg
	}
	var cmd tea.Cmd
	m.stack[len(m.stack)-1], cmd = m.Top().Update(msg)
	return m, cmd
}

func (m NavStack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.UpdateModel(msg)
}

func (m NavStack) View() string {
	return m.Top().View()
}
