package UITable

import (
	"errors"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yupanquiah/projects/task-tracker/internal/ui"
	"github.com/yupanquiah/projects/task-tracker/internal/utils"
)

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			selected := m.table.SelectedRow()
			prefix := ui.PrefixDetail.Render("Details of task")
			taskNum := ui.TaskNumber.Render(selected[0])
			return m, tea.Printf("%s %s: %s", prefix, taskNum, selected[1])
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return ui.TBaseStyle.Render(m.table.View())
}

func Create(columns []table.Column, rows []table.Row) {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	styles := table.DefaultStyles()
	styles.Header = ui.THeader
	t.SetStyles(styles)
	m := model{t}

	logger := utils.NewLogger()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		err = errors.New("error running program")
		logger.Error(err.Error())
	}
}
