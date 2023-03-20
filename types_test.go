package nav_test

import tea "github.com/charmbracelet/bubbletea"

type emptyModel struct{}

func (emptyModel) Init() tea.Cmd                       { return nil }
func (emptyModel) Update(tea.Msg) (tea.Model, tea.Cmd) { return nil, nil }
func (emptyModel) View() string                        { return "" }

type testModel struct {
	id   string
	msgs []tea.Msg
	view string
}

func initialTestModel(id string) testModel {
	return testModel{
		id:   id,
		msgs: []tea.Msg{},
	}
}

func (testModel) Init() tea.Cmd {
	return func() tea.Msg {
		return "Init"
	}
}
func (m testModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.msgs = append(m.msgs, msg)
	return m, nil
}
func (m testModel) View() string { return m.view }
