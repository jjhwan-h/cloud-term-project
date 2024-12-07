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
	createImage        option = "11"
	deleteImage        option = "12"
	quit               option = "99"
)

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	menuColumns = []table.Column{
		{Title: "number", Width: 8},
		{Title: "options", Width: 22},
		{Title: "note", Width: 45},
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
		{"1", "list instances", "List instances with a specific state"},
		{"2", "available zones", "List all available zones"},
		{"3", "start instance", "Start an instance in a stopped state"},
		{"4", "available regions", "List all available regions"},
		{"5", "stop instance", "Stop an instance in a running state"},
		{"6", "create instance", "Select an image and a security group"},
		{"7", "reboot instance", "Restart an instance in a running state"},
		{"8", "list images", "List all available images"},
		{"9", "connect instance", "Access the instance through SSH"},
		{"10", "list security groups", "List all security groups"},
		{"11", "create image", "Select a stopped or running instance"},
		{"12", "delete image", "Delete an image"},
		{"99", "quit", "..."},
	}
)
