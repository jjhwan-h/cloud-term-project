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
		case "ctrl+c":
			return cli, tea.Quit
		case "m", "M":
			cli.page = 0
			cli.menu = option(main)
			cli.ch = []string{}
			cli.table = NewTable(menuColumns, menuRows)
		case "enter":
			if cli.menu == main { //main menu
				selected := option(cli.table.SelectedRow()[0])
				if selected == startInstance { //stopped 상태인 instance만 출력
					cli.updateStoppedInstance(selected)
				} else if selected == stopInstance ||
					selected == rebootInstance ||
					selected == connectInstance { //runing 상태인 instance만 출력
					cli.updateRunningInstance(selected)
				} else if selected == createInstance ||
					selected == deleteImage { // 사용가능한 image출력
					cli.updateListImage(selected)
				} else if selected == createImage ||
					selected == terminsateInstance {
					cli.updateRnSInstance(selected)
				} else {
					cli.processAnswer(selected)
				}
			} else if cli.menu == createInstance {
				switch cli.page {
				case 1:
					cli.ch = append(cli.ch, cli.table.SelectedRow()[0])
					cli.updateListSg(cli.menu)
				case 2:
					cli.ch = append(cli.ch, cli.table.SelectedRow()[0])
					cli.processAnswer(cli.menu)
					cli.isEnd = true
				}
			} else if cli.menu == connectInstance {
				switch cli.page {
				case 1:
					cli.ch = append(cli.ch, cli.table.SelectedRow()[6])
					cli.processAnswer(cli.menu)
					cli.shell.menu = cli.menu
					cli.shell.host = cli.ch[0]
					cli.shell.Start()
					cli.isEnd = true
					return cli, tea.Batch(tea.ClearScreen)
				}
			} else if cli.menu == createImage {
				switch cli.page {
				case 1:
					cli.ch = append(cli.ch, cli.table.SelectedRow()[0])
					cli.shell = NewShell(96, 0, "Please enter an image name with at least 3 characters")
					cli.shell.menu = cli.menu
					cli.shell.Start()
					cli.ch = append(cli.ch, cli.shell.messages[0])
					cli.shell.messages = []string{}
					cli.processAnswer(cli.menu)
					return cli, tea.Batch(tea.ClearScreen)
				}
			} else if cli.menu == deleteImage {
				switch cli.page {
				case 1:
					cli.ch = append(cli.ch, cli.table.SelectedRow()[0])
					cli.processAnswer(cli.menu)
					cli.isEnd = true
				}
			} else if cli.menu == terminsateInstance {
				switch cli.page {
				case 1:
					cli.ch = append(cli.ch, cli.table.SelectedRow()[0])
					cli.ch = append(cli.ch, cli.table.SelectedRow()[5])
					cli.processAnswer(cli.menu)
					cli.isEnd = true
				}
			} else if cli.menu == startInstance ||
				cli.menu == stopInstance ||
				cli.menu == rebootInstance {
				cli.ch = append(cli.ch, cli.table.SelectedRow()[0])
				cli.processAnswer(cli.menu)
				cli.isEnd = true
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
	if len(cli.ch) > 0 {
		str := baseStyle.Render(cli.table.View()) + "\n" + cli.table.HelpView() + "\n" + cli.ch[len(cli.ch)-1]
		if cli.isEnd {
			cli.ch = []string{}
			cli.isEnd = false
		}
		return str
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
