package edit

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/emmd474/devlog/internal/model"
	"github.com/emmd474/devlog/internal/storage"
)

type editModel struct {
	entry    model.Entry
	textarea textarea.Model
}

func NewEditModel(entry model.Entry) editModel {
	ta := textarea.New()
	ta.SetValue(entry.Message)
	ta.Focus()

	return editModel{
		entry:    entry,
		textarea: ta,
	}
}

func (m editModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m editModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+s":
			m.entry.Message = m.textarea.Value()
			_ = storage.UpdateEntry(m.entry)
			return m, tea.Quit

		case "esc":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m editModel) View() string {
	return "Edit entry:\n\n" + m.textarea.View() +
		"\n\nCtrl+S to save • Esc to cancel"
}
