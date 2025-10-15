package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	Cmd = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c27aff")).
		Bold(true)

	Success = lipgloss.NewStyle().
		SetString("SUCCESS:").
		Foreground(lipgloss.Color("#26c941")).
		Bold(true)

	Warning = lipgloss.NewStyle().
		SetString("WARNING:").
		Foreground(lipgloss.Color("#ffbc30")).
		Bold(true)

	Error = lipgloss.NewStyle().
		SetString("ERROR").
		Foreground(lipgloss.Color("#FF5F87")).
		Bold(true)

	Info = lipgloss.NewStyle().
		SetString("INFO").
		Foreground(lipgloss.Color("#00dfff")).
		Bold(true)

	Id = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66"))

	DataTime = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("#6B7280"))

	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#F472B6"))

	THeader = table.DefaultStyles().Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6B7280")).
		BorderBottom(true).
		Bold(true)

	TBaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#6B7280"))

	PrefixDetail = lipgloss.NewStyle().
			Bold(true)

	TaskNumber = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66"))
)
