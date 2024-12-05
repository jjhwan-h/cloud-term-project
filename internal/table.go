package internal

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type option string

func (cli *Cli) Init() tea.Cmd { return tea.Batch(tea.EnterAltScreen, tea.EnableMouseCellMotion) }

func (cli *Cli) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return cli, tea.Quit
		case "m", "M":
			cli.menu = option(main)
			cli.table = NewTable(menuColumns, menuRows)
		case "enter":
			if cli.menu == main { //main menu
				selected := option(cli.table.SelectedRow()[0])
				if selected == startInstance { //stopped 상태인 instance만 출력
					cli.updateStoppedInstance(selected)
				} else if selected == stopInstance || selected == rebootInstance { //runing 상태인 instance만 출력
					cli.updateRunningInstance(selected)
				} else if selected == createInstance { // 사용가능한 image출력
					cli.updateListImage(selected)
				} else {
					cli.processAnswer(selected)
				}
			} else {
				cli.ch = ptr(cli.table.SelectedRow()[0])
				cli.processAnswer(cli.menu)
			}
		}
	case tea.MouseMsg:
		if tea.MouseAction(msg.Action) == tea.MouseAction(tea.MouseButtonLeft) {
			row, ok := cli.getRowFromMouse(msg.Y)
			if ok {
				cli.table.SetCursor(*row)
				key := tea.KeyMsg{Type: tea.KeyEnter}
				cli.Update(key)
			}
		}
	}

	cli.table, cmd = cli.table.Update(msg)
	return cli, cmd
}

func (cli *Cli) View() string {
	if cli.ch != nil {
		return baseStyle.Render(cli.table.View()) + "\n" + cli.table.HelpView() + "\n" + *cli.ch
	}
	return baseStyle.Render(cli.table.View()) + "\n" + cli.table.HelpView()
}

func (cli *Cli) getRowFromMouse(y int) (*int, bool) {
	rowIndex := y - 3
	if rowIndex < 0 || rowIndex >= len(cli.table.Rows()) {
		return nil, false
	}
	return &rowIndex, true
}

func NewTable(columns []table.Column, rows []table.Row) table.Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
		table.WithKeyMap(NewKeyMap()),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	t.Help.ShowAll = true
	return t
}
