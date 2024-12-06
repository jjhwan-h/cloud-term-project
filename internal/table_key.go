package internal

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
)

type KeyMap struct {
	LineUp   key.Binding
	LineDown key.Binding
	GotoTop  key.Binding // menu
}

func NewKeyMap() table.KeyMap {
	return table.KeyMap{
		LineUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		GotoTop: key.NewBinding(
			key.WithKeys("m", "M"),
			key.WithHelp("m/M", "go to menu"),
		),
	}
}
