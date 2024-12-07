package internal

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

const (
	main               option = "0"
	listInstance       option = "1"
	availableZones     option = "2"
	startInstance      option = "3"
	availableRegions   option = "4"
	stopInstance       option = "5"
	createInstance     option = "6"
	rebootInstance     option = "7"
	listImages         option = "8"
	connectInstance    option = "9"
	listSecurityGroups option = "10"
	quit               option = "99"
)

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	menuColumns = []table.Column{
		{Title: "number", Width: 8},
		{Title: "options", Width: 25},
		{Title: "note", Width: 35},
	}
	instanceColumns = []table.Column{
		{Title: "id", Width: 22},
		{Title: "AMI", Width: 22},
		{Title: "type", Width: 15},
		{Title: "state", Width: 15},
		{Title: "monitoring state", Width: 22},
		{Title: "public addr", Width: 0},
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
	sgColumns = []table.Column{
		{Title: "id", Width: 15},
		{Title: "name", Width: 25},
	}
	menuRows = []table.Row{
		{"1", "list instance", "Print all states of an instance"},
		{"2", "available zones", "..."},
		{"3", "start instance", "Print instances in a stopped state"},
		{"4", "available regions", "..."},
		{"5", "stop instance", "Print instances in a running state"},
		{"6", "create instance", "..."},
		{"7", "reboot instance", "Print instances in a running state"},
		{"8", "list images", "Print all available images"},
		{"9", "connect instance", "Access the instance through SSH"},
		{"10", "list security groups", "..."},
		{"99", "quit", "..."},
	}
)
