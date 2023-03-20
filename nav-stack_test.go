package nav_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	nav "github.com/vknabel/go-bubblenav"
)

func TestNavStackDisplaysPushed(t *testing.T) {
	prog := tea.NewProgram(
		nav.NewStack(initialTestModel("initial")),
		tea.WithoutCatchPanics(),
		tea.WithoutRenderer(),
	)
	go func() {
		prog.Send(nav.PagePushMsg{initialTestModel("pushed")})
		prog.Quit()
	}()
	model, err := prog.StartReturningModel()
	if err != nil {
		t.Error(err)
	}
	navPage, ok := model.(nav.NavStack)
	if !ok {
		t.Error("expected model to be a NavPage, got", model)
	}
	testModel, ok := navPage.Top().(testModel)
	if !ok {
		t.Error("expected model to be a testModel, got", model)
	}

	if testModel.id != "pushed" {
		t.Error("expected model id to be 'pushed', got", testModel.id)
	}
}

func TestNavStackDisplaysInitialWhenPushedAndPopped(t *testing.T) {
	prog := tea.NewProgram(
		nav.NewStack(initialTestModel("initial")),
		tea.WithoutCatchPanics(),
		tea.WithoutRenderer(),
	)
	go func() {
		prog.Send(nav.PagePushMsg{initialTestModel("pushed")})
		prog.Send(nav.PagePopMsg{})
		prog.Quit()
	}()
	model, err := prog.StartReturningModel()
	if err != nil {
		t.Error(err)
	}
	navPage, ok := model.(nav.NavStack)
	if !ok {
		t.Error("expected model to be a NavPage, got", model)
	}
	testModel, ok := navPage.Top().(testModel)
	if !ok {
		t.Error("expected model to be a testModel, got", model)
	}

	if testModel.id != "initial" {
		t.Error("expected model id to be 'initial', got", testModel.id)
	}
}
