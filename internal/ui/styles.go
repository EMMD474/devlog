package ui

import "github.com/charmbracelet/lipgloss"

var (
	HeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Padding(0, 1)

	DateStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C7086"))

	MessageStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#CDD6F4"))

	BoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Padding(1, 2)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A6E3A1"))

	BulletStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F9E2AF"))

	TimeStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#89B4FA"))

	EmptyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C7086")).
		Italic(true)
)
