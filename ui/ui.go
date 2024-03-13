package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DeepAung/rubik/rubik"
	"github.com/DeepAung/rubik/ui/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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
	red      = lipgloss.Color("#FF0000")
	hotPink  = lipgloss.Color("#FF06B7")
	darkPink = lipgloss.Color("#79305a")
	darkGray = lipgloss.Color("#767676")
)

type myKeyMap struct {
	Next  key.Binding
	Prev  key.Binding
	Esc   key.Binding
	Enter key.Binding
	Help  key.Binding
	Quit  key.Binding
}

var keys = myKeyMap{
	Next: key.NewBinding(
		key.WithKeys("tab", "down"),
		key.WithHelp("tab/↓", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("shift+tab", "up"),
		key.WithHelp("shift+tab/↑", "previous"),
	),
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "unfocus"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "focus or execute function"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
}

func (k myKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}
func (k myKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Next, k.Prev, k.Enter},
		{k.Help, k.Quit},
	}
}

type model struct {
	rubik rubik.IRubik

	// lipgloss' components
	rotateInput textinput.Model
	cycleInput  textinput.Model
	help        help.Model

	// states
	focusIdx     int
	currentMoves int
	currentTimes int
	errorText    string

	// info
	fullWidth  int
	fullHeight int
	keys       help.KeyMap
}

func initialModel() model {
	fullWidth, fullHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	rotateInput := textinput.New()
	rotateInput.Focus()
	rotateInput.Placeholder = "notations..."
	rotateInput.CharLimit = 156
	rotateInput.Width = 20

	cycleInput := textinput.New()
	cycleInput.Placeholder = "notations..."
	cycleInput.CharLimit = 156
	cycleInput.Width = 20

	return model{
		rubik: rubik.NewRubik(),

		rotateInput: rotateInput,
		cycleInput:  cycleInput,
		help:        help.New(),

		focusIdx:     0,
		currentMoves: 0,
		currentTimes: 0,
		errorText:    "",

		fullWidth:  fullWidth,
		fullHeight: fullHeight,
		keys:       keys,
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

		case "?":
			m.help.ShowAll = !m.help.ShowAll
			return m, nil

		case "tab", "shift+tab", "up", "down":
			s := msg.String()

			if s == "tab" || s == "down" {
				m.focusIdx = (m.focusIdx + 1) % 5
			} else {
				m.focusIdx = (m.focusIdx - 1 + 5) % 5
			}

			switch m.focusIdx {
			case 0:
				m.cycleInput.Blur()
				return m, m.rotateInput.Focus()
			case 2:
				m.rotateInput.Blur()
				return m, m.cycleInput.Focus()
			default:
				m.rotateInput.Blur()
				m.cycleInput.Blur()
				return m, nil
			}

		case "esc":
			switch m.focusIdx {
			case 0:
				m.rotateInput.Blur()
				return m, nil
			case 2:
				m.cycleInput.Blur()
				return m, nil
			}

		case "enter":
			m.errorText = ""

			switch m.focusIdx {

			case 0:
				if !m.rotateInput.Focused() {
					m.rotateInput.Focus()
					return m, nil
				}

				_, err := m.rubik.Rotates(m.rotateInput.Value(), true)
				if err != nil {
					m.errorText = err.Error()
					return m, nil
				}

			case 1:
				m.rubik.Reset(true)

			case 2:
				if !m.cycleInput.Focused() {
					m.cycleInput.Focus()
					return m, nil
				}

				times, moves, err := m.rubik.CycleNumber(m.cycleInput.Value())
				if err != nil {
					m.errorText = err.Error()
					return m, nil
				}

				m.currentTimes = times
				m.currentMoves = moves

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
		"",
		m.ErrorView(),
		"",
		m.help.View(m.keys),
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

// TODO: create rubik view using lipgloss
func (m model) RubikView() string {
	return lipgloss.NewStyle().
		Width(m.fullWidth/2-2).
		Align(lipgloss.Center, lipgloss.Center).
		Render(
			lipgloss.NewStyle().
				BorderStyle(lipgloss.HiddenBorder()).
				Render(m.rubik.Sprint()),
		)
}

func (m model) ActionsView() string {
	sections := [][]string{
		{"[Rotate]", m.rotateInput.View(), ""},
		{"[Reset]\n"},
		{
			"[CycleNumber]",
			m.cycleInput.View(),
			fmt.Sprintf("times: %d\tmoves: %d\n", m.currentTimes, m.currentMoves),
		},
		{"[Undo]\n"},
		{"[Redo]\n"},
	}

	canUndo := m.rubik.CanUndo()
	canRedo := m.rubik.CanRedo()

	if !canUndo {
		if m.focusIdx == 3 {
			utils.SetColor(&sections[3][0], darkPink)
		} else {
			utils.SetColor(&sections[3][0], darkGray)
		}
	}

	if !canRedo {
		if m.focusIdx == 4 {
			utils.SetColor(&sections[4][0], darkPink)
		} else {
			utils.SetColor(&sections[4][0], darkGray)
		}
	}

	// cannot override the previous SetColor()
	utils.SetColor(&sections[m.focusIdx][0], hotPink)

	result := make([]string, len(sections))
	for idx, section := range sections {
		result[idx] = strings.Join(section, "\n")
	}

	return lipgloss.NewStyle().
		Width(m.fullWidth / 2).
		Render(lipgloss.JoinVertical(lipgloss.Left, result...))
}

func (m model) ErrorView() string {
	if m.errorText == "" {
		return ""
	}

	return lipgloss.NewStyle().Foreground(red).Italic(true).Render("error: ", m.errorText)
}
