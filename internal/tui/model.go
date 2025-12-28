package tui

import (
	"fmt"

	"github.com/atomisadev/portkey/internal/config"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	host config.Host
}

func (i item) Title() string       { return i.host.Alias }
func (i item) Description() string { return fmt.Sprintf("%s (%s)", i.host.Hostname, i.host.User) }
func (i item) FilterValue() string { return i.host.Alias + " " + i.host.Hostname }

type model struct {
	list   list.Model
	width  int
	height int
	loaded bool
}

func initialModel() model {
	return model{
		list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
	}
}

type hostsLoadedMsg []config.Host
type errMsg error

func loadHostsCmd() tea.Cmd {
	return func() tea.Msg {
		hosts, err := config.LoadHosts()
		if err != nil {
			return errMsg(err)
		}
		return hostsLoadedMsg(hosts)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.SetWindowTitle("PortKey"), loadHostsCmd())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		h, v := AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case hostsLoadedMsg:
		m.loaded = true
		var items []list.Item
		for _, h := range msg {
			items = append(items, item{host: h})
		}

		delegate := list.NewDefaultDelegate()
		delegate.Styles.SelectedTitle = SelectedTitle
		delegate.Styles.SelectedDesc = SelectedDesc

		m.list = list.New(items, delegate, m.width, m.height)
		m.list.Title = "SSH Connections"
		m.list.SetShowHelp(false)
		m.list.Styles.Title = TitleStyle

		h, v := AppStyle.GetFrameSize()
		m.list.SetSize(m.width-h, m.height-v)

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading configuration..."
	}

	return AppStyle.Render(m.list.View())
}

func Start() error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
