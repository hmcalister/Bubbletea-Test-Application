package model

import "github.com/charmbracelet/lipgloss"

var (
	color = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA"))
		// Background(lipgloss.Color("#7D56F4"))

	headerStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Width(50).
			Inherit(color)

	listItemBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4"))

	textInputStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true, false, false, false).
			Width(50).
			BorderForeground(lipgloss.Color("#7D56F4"))

	activeListItemStyle = lipgloss.NewStyle().
				Bold(true).
				Inherit(color)

	inactiveListItemStyle = lipgloss.NewStyle().
				Faint(true).
				Inherit(color)
)
