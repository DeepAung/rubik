package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/DeepAung/rubik/rubik"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

type model struct {
	rubik       rubik.IRubik
	rotateInput textinput.Model
	cycleInput  textinput.Model
	focusIdx    int
	fullWidth   int
	fullHeight  int
}

func initialModel() model {
	fullWidth, fullHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	rotateInput := textinput.New()
	rotateInput.Placeholder = "notations..."
	rotateInput.CharLimit = 156
	rotateInput.Width = 20

	cycleInput := textinput.New()
	cycleInput.Placeholder = "notations..."
	cycleInput.CharLimit = 156
	cycleInput.Width = 20

	return model{
		rubik:       rubik.NewRubik(),
		rotateInput: rotateInput,
		cycleInput:  cycleInput,
		focusIdx:    0,
		fullWidth:   fullWidth,
		fullHeight:  fullHeight,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab", "shift+tab":
			s := msg.String()

			if s == "tab" {
				m.focusIdx = (m.focusIdx + 1) % 5
			} else {
				m.focusIdx = (m.focusIdx - 1 + 5) % 5
			}

			if m.focusIdx == 0 {
				m.cycleInput.Blur()
				return m, m.rotateInput.Focus()
			} else if m.focusIdx == 2 {
				m.rotateInput.Blur()
				return m, m.cycleInput.Focus()
			} else {
				m.rotateInput.Blur()
				m.cycleInput.Blur()
				return m, nil
			}

		case "enter":
			switch m.focusIdx {
			case 0:
				_, _ = m.rubik.Rotates(m.rotateInput.Value(), true)
			case 1:
				m.rubik.Reset(true)
			case 2:
				_, _, _ = m.rubik.CycleNumber(m.cycleInput.Value())
			case 3:
				m.rubik.Undo(1)
			case 4:
				m.rubik.Redo(1)
			}

		}
	}

	var cmd1, cmd2 tea.Cmd
	m.rotateInput, cmd1 = m.rotateInput.Update(msg)
	m.cycleInput, cmd2 = m.cycleInput.Update(msg)

	return m, tea.Batch(cmd1, cmd2)
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		m.HeaderView(),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.RubikView(),
			m.ActionsView(),
		),
	)
}

func (m model) HeaderView() string {
	return lipgloss.NewStyle().
		Width(m.fullWidth).
		Height(3).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Render("Rubik Simulator")
}

func (m model) RubikView() string {
	return lipgloss.NewStyle().
		Width(m.fullWidth / 2).
		Render(m.rubik.Sprint())
}

func (m model) ActionsView() string {
	sections := []string{
		"[Rotate]\n" + m.rotateInput.View() + "\n",
		"[Reset]\n",
		"[CycleNumber]\n" + m.cycleInput.View() + "\n",
		"[Undo]\n",
		"[Redo]\n",
	}

	sections[m.focusIdx] = lipgloss.
		NewStyle().
		Foreground(hotPink).
		Render(sections[m.focusIdx])

	return lipgloss.NewStyle().
		Width(m.fullWidth / 2).
		Render(lipgloss.JoinVertical(lipgloss.Left, sections...))
}
