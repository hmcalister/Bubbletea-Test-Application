package model

import "github.com/charmbracelet/lipgloss"

var (
	colorWhite  = lipgloss.Color("#FAFAFA")
	colorPurple = lipgloss.Color("#7D56F4")

	textColor = lipgloss.NewStyle().
			Foreground(colorWhite)

	borderColor = lipgloss.NewStyle().
			BorderForeground(colorPurple)

	headerStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true, false, true, false).
			Width(80).
			Inherit(textColor).
			Inherit(borderColor)

	listItemBorder = lipgloss.NewStyle().
			PaddingRight(2).
			Border(lipgloss.RoundedBorder()).
			Inherit(borderColor)

	textInputStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true, false, true, false).
			Width(80).
			Inherit(borderColor)

	activeListItemStyle = lipgloss.NewStyle().
				Bold(true).
				Inherit(textColor)

	inactiveListItemStyle = lipgloss.NewStyle().
				Faint(true).
				Inherit(textColor)

	paginatorActiveDotStyle = lipgloss.NewStyle().
				Foreground(colorWhite)

	paginatorInactiveDotStyle = lipgloss.NewStyle().
					Foreground(colorPurple)
)
