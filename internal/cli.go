package internal

import (
	"TermProject/aws"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Cli struct {
	aws   aws.Aws
	table table.Model
	shell *shell
	ch    []string
	menu  option
	page  int
	isEnd bool
}

func NewCli() (*Cli, error) {
	aws, err := aws.NewAws()
	if err != nil {
		return nil, err
	}
	t := NewTable(menuColumns, menuRows)
	s := NewShell(100, 25, "Send a command...(ESC/Ctrl+c exit)")
	return &Cli{
		aws:   *aws,
		table: t,
		shell: s,
		ch:    nil,
		menu:  option(main),
		page:  0,
		isEnd: false,
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
		handleResult(nil, err)
		cli.table = NewTable(instanceColumns, rows)
	case availableZones:
		rows, err := cli.aws.AvailableZones()
		handleResult(nil, err)
		cli.table = NewTable(zoneColumns, rows)
	case startInstance:
		res, err := cli.aws.StartInstance(cli.ch)
		cli.ch = append(cli.ch, *handleResult(res, err))
	case availableRegions:
		rows, err := cli.aws.AvailableRegions()
		handleResult(nil, err)
		cli.table = NewTable(regionColumns, rows)
	case stopInstance:
		res, err := cli.aws.StopInstance(cli.ch)
		cli.ch = append(cli.ch, *handleResult(res, err))
	case createInstance:
		res, err := cli.aws.CreateInstance(cli.ch)
		cli.ch = append(cli.ch, *handleResult(res, err))
	case rebootInstance:
		res, err := cli.aws.RebootInstance(cli.ch)
		cli.ch = append(cli.ch, *handleResult(res, err))
	case listImages:
		rows, err := cli.aws.ListImages()
		handleResult(nil, err)
		cli.table = NewTable(imageColumns, rows)
	case connectInstance:
		conn, err := cli.aws.ConnectInstance(cli.ch)
		if err != nil {
			log.Println(err)
			return
		}
		cli.shell.conn = conn
	case listSecurityGroups:
		rows, err := cli.aws.ListSecurityGroup()
		handleResult(nil, err)
		cli.table = NewTable(sgColumns, rows)
	case createImage:
		res, err := cli.aws.CreateImage(cli.ch)
		cli.ch = append(cli.ch, *handleResult(res, err))
	case deleteImage:
		res, err := cli.aws.DeleteImage(cli.ch)
		cli.ch = append(cli.ch, *handleResult(res, err))
	case terminsateInstance:
		res, err := cli.aws.TerminateInstance(cli.ch)
		cli.ch = append(cli.ch, *handleResult(res, err))
	case quit:
		os.Exit(0)
	}
}

func (cli *Cli) updateRunningInstance(selected option) {
	rows, err := cli.aws.ListInstances(ptr("running"))
	handleResult(nil, err)
	cli.table = NewTable(instanceColumns, rows)
	cli.menu = selected
	cli.page++
}
func (cli *Cli) updateStoppedInstance(selected option) {
	rows, err := cli.aws.ListInstances(ptr("stopped"))
	handleResult(nil, err)
	cli.table = NewTable(instanceColumns, rows)
	cli.menu = selected
	cli.page++
}
func (cli *Cli) updateRnSInstance(selected option) { // Running & Stopped
	rows, err := cli.aws.ListInstances(ptr("running"))
	handleResult(nil, err)
	rows2, err := cli.aws.ListInstances(ptr("stopped"))
	handleResult(nil, err)
	rows = append(rows, rows2...)
	cli.table = NewTable(instanceColumns, rows)
	cli.menu = selected
	cli.page++
}
func (cli *Cli) updateListImage(selected option) {
	rows, err := cli.aws.ListImages()
	handleResult(nil, err)
	cli.table = NewTable(imageColumns, rows)
	cli.menu = selected
	cli.page++
}
func (cli *Cli) updateListSg(selected option) {
	rows, err := cli.aws.ListSecurityGroup()
	handleResult(nil, err)
	cli.table = NewTable(sgColumns, rows)
	cli.menu = selected
	cli.page++
}
