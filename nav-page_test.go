package nav_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	nav "github.com/vknabel/go-bubblenav"
)

func TestNavPageDisplaysPushed(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer

	prog := tea.NewProgram(
		nav.NewPage(initialTestModel("initial")),
		tea.WithoutCatchPanics(),
		tea.WithoutRenderer(),
		tea.WithInput(&in),
		tea.WithOutput(&out),
	)
	go func() {
		prog.Send(nav.PagePushMsg{initialTestModel("pushed")})
		go prog.Quit()
	}()
	model, err := prog.StartReturningModel()
	if err != nil {
		t.Error(err)
	}
	stack, ok := model.(nav.NavPage)
	if !ok {
		t.Error("expected model to be a NavPage, got", model)
	}
	testModel, ok := stack.Top().(testModel)
	if !ok {
		t.Error("expected model to be a testModel, got", model)
	}

	if testModel.id != "pushed" {
		t.Error("expected model id to be 'pushed', got", testModel.id)
	}
}

func TestNavPageDisplaysInitialWhenPushedAndPopped(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer

	prog := tea.NewProgram(
		nav.NewPage(initialTestModel("initial")),
		tea.WithoutCatchPanics(),
		tea.WithoutRenderer(),
		tea.WithInput(&in),
		tea.WithOutput(&out),
	)
	go func() {
		prog.Send(nav.PagePushMsg{initialTestModel("pushed")})
		go func() {
			prog.Send(nav.PagePopMsg{})
			go prog.Quit()
		}()
	}()
	model, err := prog.StartReturningModel()
	if err != nil {
		t.Error(err)
	}
	stack, ok := model.(nav.NavPage)
	if !ok {
		t.Error("expected model to be a NavPage, got", model)
	}
	testModel, ok := stack.Top().(testModel)
	if !ok {
		t.Error("expected model to be a testModel, got", model)
	}

	if testModel.id != "initial" {
		t.Error("expected model id to be 'initial', got", testModel.id)
	}
}
