package edit

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/emmd474/devlog/internal/model"
)

var listStyle = lipgloss.NewStyle().Margin(1, 2)

type selectModel struct {
	list    list.Model
	entries []model.Entry
	width   int
	height  int
}

func NewSelectModel(entries []model.Entry) selectModel {
	items := make([]list.Item, len(entries))
	for i, e := range entries {
		items[i] = item{entry: e}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select a log to edit"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)

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

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h, v := listStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.list.FilterState() == list.Filtering {
				break
			}
			selected := m.list.SelectedItem().(item).entry
			return NewEditModel(selected, m.width, m.height), nil
		case "q", "ctrl+c":
			if m.list.FilterState() == list.Filtering {
				break
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m selectModel) View() string {
	return listStyle.Render(m.list.View())
}
