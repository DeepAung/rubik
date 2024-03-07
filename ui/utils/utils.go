package utils

import "github.com/charmbracelet/lipgloss"

func SetColor(str *string, color lipgloss.TerminalColor) {
	*str = lipgloss.NewStyle().UnsetForeground().Foreground(color).Render(*str)
}
