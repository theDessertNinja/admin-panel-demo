package button

import (
	"math/rand/v2"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	label string
	id    int

	ButtonBlurStyle  lipgloss.Style
	ButtonFocusStyle lipgloss.Style
	LabelBlurStyle   lipgloss.Style
	LabelFocusStyle  lipgloss.Style

	width  int
	height int

	keyMap keyMap

	focus bool
}

type ButtonMsg struct {
	Id int
}

type keyMap struct {
	enter key.Binding
}

func New() Model {
	defaultKeyMap := keyMap{
		enter: key.NewBinding(key.WithKeys("enter")),
	}

	focusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // Black text
		Background(lipgloss.Color("7")). // White background
		Bold(true).
		Padding(0, 2).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center). // Center text within the button
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("7"))

	blurStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")). // White text
		Background(lipgloss.Color("8")). // Bright black (gray) background
		Padding(0, 2).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center). // Center text within the button
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8"))

	model := Model{
		label:            "Unnamed Button",
		id:               rand.Int(),
		ButtonBlurStyle:  blurStyle,
		ButtonFocusStyle: focusStyle,
		LabelBlurStyle:   lipgloss.NewStyle(),
		LabelFocusStyle:  lipgloss.NewStyle(),
		width:            20,
		height:           3,
		keyMap:           defaultKeyMap,
		focus:            false,
	}

	return model
}

func (m *Model) SetLabel(label string) {
	m.label = label
}

func (m Model) GetLabel() string {
	return m.label
}

func (m *Model) SetId(id int) {
	m.id = id
}

func (m Model) GetId() int {
	return m.id
}

func (m *Model) SetWidth(width int) {
	m.width = width
}

func (m Model) Width() int {
	return m.width
}

func (m *Model) SetHeight(height int) {
	m.height = height
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if m.focus == false {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.enter):
			return m, func() tea.Msg { return ButtonMsg{Id: m.id} }
		}
	}

	return m, nil
}

func (m Model) View() string {
	var labelStyle, buttonStyle lipgloss.Style
	if m.focus {
		labelStyle = m.LabelFocusStyle
		buttonStyle = m.ButtonFocusStyle
	} else {
		labelStyle = m.LabelBlurStyle
		buttonStyle = m.ButtonBlurStyle
	}

	labelRender := labelStyle.Render(m.label)
	buttonRender := buttonStyle.Width(m.width).Height(m.height).Render(labelRender)

	return buttonRender
}
