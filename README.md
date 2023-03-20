# go-bubblenav

The goal of bubblenav is to provide easy to use navigation helpers for [bubbletea](https://github.com/charmbracelet/bubbletea) applications.

Simply wrap your actual models with `nav.NavPage` or `nav.NavStack` and emit `nav.Push(tea.Model) tea.Cmd` to push a new app model as a new page. Call `nav.Pop()` to remove the current page from the top of the stack.

- Bubblenav will call the `Init()` method of your `tea.Model`.
- Only the currently active page model will receive `Update`.
- Recorded `tea.WindowSizeMsg` will be repeated for new pages.
- `esc` will pop the currnet page.
- Unless `DisableQuitKeyBindings()` called, `q` and `ctrl+c` will quit.
- With `Replace` you navigate irreversibly and replace the current page.

```go
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	nav "github.com/vknabel/go-bubblenav"
)

func main() {
    var initial tea.Model = initialModel()
    wrapped := nav.NavPage(initial)
    err := tea.NewProgram(wrapped).Start()
    if err != nil {
        panic(err)
    }
}

// ...

func (m myModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "n":
            return m, nav.Push(nextPage())
        }
    }
    return m, nil
}
```

## Installation

```bash
go get github.com/vknabel/go-bubblenav
```

## Future Development

- [ ] Decide between `nav.NavPage` and `nav.NavStack` and remove the other. _Both are different implementations of the same concept._
- [ ] Implement a modal.

# License

Licensed under the [MIT](./LICENSE) License.
