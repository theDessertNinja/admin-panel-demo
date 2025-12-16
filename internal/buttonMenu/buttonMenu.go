package buttonmenu

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thedessertninja/admin-panel/internal/button"
)

type Model struct {
	title string

	selections       []button.Model
	currentSelection int

	TitleStyle lipgloss.Style
	MenuStyle  lipgloss.Style

	keyMap keyMap
}

type keyMap struct {
	up   key.Binding
	down key.Binding
}

type MenuMsg struct {
	id int
}

var (
	defaultTitleStyle = lipgloss.NewStyle()
	defaultMenuStyle  = lipgloss.NewStyle()
	defaultKeyMap     = keyMap{
		up:   key.NewBinding(key.WithKeys("up", "k")),
		down: key.NewBinding(key.WithKeys("down", "j")),
	}
)

func New() Model {
	model := Model{
		title:            "Menu",
		selections:       []button.Model{},
		currentSelection: 0,
		TitleStyle:       defaultTitleStyle,
		MenuStyle:        defaultMenuStyle,
		keyMap:           defaultKeyMap,
	}

	return model
}

func (m *Model) SetTitle(title string) {
	m.title = title
}

func (m Model) GetTitle() string {
	return m.title
}

func (m *Model) FromLabels(labels []string) {
	for i, label := range labels {
		button := button.New()
		button.SetLabel(label)
		button.SetId(i)
		if len(m.selections) == i {
			m.selections = append(m.selections, button)
		} else {
			m.selections[i] = button
		}
	}
}

func (m *Model) IncrementSelection() {
	if m.currentSelection == len(m.selections)-1 {
		m.currentSelection = 0
	} else {
		m.currentSelection++
	}
}

func (m *Model) DecrementSelection() {
	if m.currentSelection == 0 {
		m.currentSelection = len(m.selections) - 1
	} else {
		m.currentSelection--
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.up):
			m.DecrementSelection()
		case key.Matches(msg, m.keyMap.down):
			m.IncrementSelection()
		}
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	for i := range m.selections {
		if m.currentSelection == i {
			m.selections[i].Focus()
		} else {
			m.selections[i].Blur()
		}

		m.selections[i], cmd = m.selections[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var s string
	for _, button := range m.selections {
		s += button.View() + "\n"
	}
	return s
}
