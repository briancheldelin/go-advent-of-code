package gui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

type AdventChallange interface {
	DoWork() bool
	Render() string
	Result() int
}

type Model struct {
	title     string
	done      bool
	ready     bool
	viewport  viewport.Model
	fps       int
	challange AdventChallange
}

func NewGUI(title string, c AdventChallange, fps int) *Model {
	return &Model{
		title:     title,
		done:      false,
		fps:       fps,
		challange: c,
	}
}

// We handle the main gui loop here.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k":
			fmt.Println("Killing program")
			os.Exit(0)
		case "f":
			m.fps = m.fps * 2
		case "s":
			m.fps = m.fps / 2
		}

		// Handle keyboard and mouse events in the viewport
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)

	case frameMsg:

		if !m.ready || m.done {
			cmds = append(cmds, m.schedule())
			break
		}

		if !m.done {
			// Call the models Tick function for the frame
			// This is called at the speed of ticks
			m.done = !m.challange.DoWork()

			// Extract our grid as a color grid
			m.viewport.SetContent(m.challange.Render())

			cmds = append(cmds, m.schedule())
		}

	// Setup the initial tea window and handle resizing
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.challange.Render())
			m.ready = true
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() (s string) {
	if !m.ready {
		return "\n  Initializing..."
	}

	s += fmt.Sprintf("%s\n", m.headerView())
	s += fmt.Sprintf("%s\n", m.viewport.View())
	s += m.footerView()

	return s
}

func (m *Model) headerView() string {
	title := titleStyle.Render(fmt.Sprintf("%s, Current Result %d", m.title, m.challange.Result()))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m *Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m *Model) Init() tea.Cmd {
	return m.schedule()
}

type frameMsg struct{}

func (m *Model) schedule() tea.Cmd {
	tickInterval := time.Second / time.Duration(m.fps)
	return tea.Tick(tickInterval, func(_ time.Time) tea.Msg {
		return frameMsg{}
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
