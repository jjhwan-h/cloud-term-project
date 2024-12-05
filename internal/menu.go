package internal

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

const (
	main             option = "0"
	listInstance     option = "1"
	availableZones   option = "2"
	startInstance    option = "3"
	availableRegions option = "4"
	stopInstance     option = "5"
	createInstance   option = "6"
	rebootInstance   option = "7"
	listImages       option = "8"
	quit             option = "99"
)

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	menuColumns = []table.Column{
		{Title: "num", Width: 10},
		{Title: "options", Width: 30},
	}
	instanceColumns = []table.Column{
		{Title: "id", Width: 22},
		{Title: "AMI", Width: 22},
		{Title: "type", Width: 15},
		{Title: "state", Width: 15},
		{Title: "monitoring state", Width: 22},
	}
	imageColumns = []table.Column{
		{Title: "AMI", Width: 22},
		{Title: "name", Width: 22},
	}
	zoneColumns = []table.Column{
		{Title: "id", Width: 15},
		{Title: "region", Width: 15},
		{Title: "zone", Width: 15},
	}
	regionColumns = []table.Column{
		{Title: "region", Width: 20},
		{Title: "endpoint", Width: 35},
	}
	menuRows = []table.Row{
		{"1", "list instance"},
		{"2", "available zones"},
		{"3", "start instance"},
		{"4", "available regions"},
		{"5", "stop instance"},
		{"6", "create instance"},
		{"7", "reboot instance"},
		{"8", "list images"},
		{"99", "quit"},
	}
)
