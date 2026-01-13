package edit

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/emmd474/devlog/internal/model"
)

type selectModel struct {
	list    list.Model
	entries []model.Entry
}

func NewSelectModel(entries []model.Entry) selectModel {
	items := make([]list.Item, len(entries))
	for i, e := range entries {
		items[i] = item{entry: e}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a log to edit"

	return selectModel{
		list:    l,
		entries: entries,
	}
}

func (m selectModel) Init() tea.Cmd {
	return nil
}

func (m selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "enter" {
			selected := m.list.SelectedItem().(item).entry
			return NewEditModel(selected), nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m selectModel) View() string {
	return m.list.View()
}
