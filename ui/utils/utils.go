package utils

import (
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func SetColor(str *string, color lipgloss.TerminalColor) {
	*str = lipgloss.NewStyle().UnsetForeground().Foreground(color).Render(*str)
}

func GetFullWidthHeight() (int, int) {
	fullWidth, fullHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	return fullWidth, fullHeight
}
