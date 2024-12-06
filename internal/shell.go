package internal

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

type (
	errMsg error
)

type shell struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	host        string
	conn        *ssh.Client
	err         error
}

func NewShell() *shell {
	ta := textarea.New()
	ta.Placeholder = "Send a command...(ESC/Ctrl+c exit)"
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(100)
	ta.SetHeight(2)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(100, 25)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return &shell{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func (s *shell) Start() error {
	header := fmt.Sprintf("You are currently connected to %s\n", s.host)
	s.viewport.SetContent(header)

	p := tea.NewProgram(s)

	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (s *shell) Init() tea.Cmd {
	return textarea.Blink
}

func (s *shell) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	s.textarea, tiCmd = s.textarea.Update(msg)
	s.viewport, vpCmd = s.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			s.conn.Close()
			s.messages = []string{}
			return s, tea.Quit
		case tea.KeyEnter:
			s.updateView(s.textarea.Value())
			res, err := s.sendCmd(s.textarea.Value())
			if err != nil {
				s.updateView(err.Error())
			} else {
				s.updateView(string(res))
			}
			s.textarea.Reset()
			s.viewport.GotoBottom()
		}

	case errMsg:
		s.err = msg
		return s, nil
	}

	return s, tea.Batch(tiCmd, vpCmd)
}

func (s *shell) updateView(t string) {
	s.messages = append(s.messages, s.senderStyle.Render(viper.GetString("USER")+"$")+t)
	s.viewport.SetContent(strings.Join(s.messages, "\n"))
}
func (s *shell) sendCmd(cmd string) ([]byte, error) {
	session, err := s.newSession()
	if err != nil {
		return []byte{}, err
	}
	defer session.Close()

	cmd = strings.TrimSuffix(cmd, "\n")
	res, err := session.CombinedOutput(cmd)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}
func (s *shell) newSession() (*ssh.Session, error) {
	session, err := s.conn.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *shell) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		s.viewport.View(),
		s.textarea.View(),
	) + "\n\n"
}
