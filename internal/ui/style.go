package ui

import "github.com/charmbracelet/lipgloss"

var (
	gold   = lipgloss.Color("#FFD700")
	silver = lipgloss.Color("#CCCCCC")
	grey   = lipgloss.Color("#555555")
	red    = lipgloss.Color("#FF5F5F")
	green  = lipgloss.Color("#5FFF5F")
	onyx   = lipgloss.Color("#121212")

	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(onyx).
			Background(gold).
			Padding(0, 1).
			Bold(true)

	headerStyle = lipgloss.NewStyle().Foreground(gold)
	
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(gold).
				Bold(true).
				PaddingLeft(1)

	itemStyle = lipgloss.NewStyle().
			Foreground(silver).
			PaddingLeft(2)

	otpCodeStyle = lipgloss.NewStyle().
			Foreground(gold).
			Bold(true).
			Padding(0, 1).
			Background(lipgloss.Color("#222222"))

	helpStyle = lipgloss.NewStyle().
			Foreground(grey).
			MarginTop(1)

	progressBarStyle = lipgloss.NewStyle().
				Foreground(gold)

	errorStyle = lipgloss.NewStyle().
			Foreground(red)

	successStyle = lipgloss.NewStyle().
			Foreground(green)

	dimStyle = lipgloss.NewStyle().
			Foreground(grey)
)
