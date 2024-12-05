package internal

import (
	"TermProject/aws"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Cli struct {
	aws   aws.Aws
	table table.Model
	ch    *string
	menu  option
}

func NewCli() (*Cli, error) {
	aws, err := aws.NewAws()
	if err != nil {
		return nil, err
	}
	t := NewTable(menuColumns, menuRows)

	return &Cli{
		aws:   *aws,
		table: t,
		ch:    nil,
		menu:  option(main),
	}, nil
}

func (cli Cli) Start() error {
	if _, err := tea.NewProgram(&cli).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return nil
}

func (cli *Cli) processAnswer(choice option) {
	switch choice {
	case listInstance:
		rows, err := cli.aws.ListInstances(nil)
		cli.ch = handleResult(nil, err)
		cli.table = NewTable(instanceColumns, rows)
		cli.menu = listInstance
	case availableZones:
		rows, err := cli.aws.AvailableZones()
		cli.ch = handleResult(nil, err)
		cli.table = NewTable(zoneColumns, rows)
		cli.menu = availableZones
	case startInstance:
		res, err := cli.aws.StartInstance(*cli.ch)
		cli.ch = handleResult(res, err)
	case availableRegions:
		rows, err := cli.aws.AvailableRegions()
		cli.ch = handleResult(nil, err)
		cli.table = NewTable(regionColumns, rows)
		cli.menu = availableRegions
	case stopInstance:
		res, err := cli.aws.StopInstance(*cli.ch)
		cli.ch = handleResult(res, err)
	case createInstance:
		res, err := cli.aws.CreateInstance(*cli.ch)
		cli.ch = handleResult(res, err)
	case rebootInstance:
		res, err := cli.aws.RebootInstance(*cli.ch)
		cli.ch = handleResult(res, err)
	case listImages:
		rows, err := cli.aws.ListImages()
		cli.ch = handleResult(nil, err)
		cli.table = NewTable(imageColumns, rows)
		cli.menu = listImages
	case quit:
		os.Exit(0)
	}
}

func (cli *Cli) updateRunningInstance(selected option) {
	rows, err := cli.aws.ListInstances(ptr("running"))
	handleResult(nil, err)
	cli.table = NewTable(instanceColumns, rows)
	cli.menu = selected
}
func (cli *Cli) updateStoppedInstance(selected option) {
	rows, err := cli.aws.ListInstances(ptr("stopped"))
	handleResult(nil, err)
	cli.table = NewTable(instanceColumns, rows)
	cli.menu = selected
}
func (cli *Cli) updateListImage(selected option) {
	rows, err := cli.aws.ListImages()
	handleResult(nil, err)
	cli.table = NewTable(imageColumns, rows)
	cli.menu = selected
}
