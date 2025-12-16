package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thedessertninja/admin-panel/internal/button"
	buttonmenu "github.com/thedessertninja/admin-panel/internal/buttonMenu"
)

type model struct {
	menu buttonmenu.Model

	termW int
	termH int
}

func newModel() model {
	model := model{
		menu:  buttonmenu.New(),
		termW: 0,
		termH: 0,
	}
	model.menu.FromLabels([]string{"button 1", "button 2", "button 3"})

	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.termW = msg.Width
		m.termH = msg.Height
	case button.ButtonMsg:
		switch msg.Id {
		case 0:
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)
	return m, cmd
}

func (m model) View() string {
	menuRender := m.menu.View()
	return lipgloss.NewStyle().Width(m.termW).Height(m.termH).Align(lipgloss.Center, lipgloss.Center).Render(menuRender)
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}
