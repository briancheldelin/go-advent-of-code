package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const useHighPerformanceRenderer = false

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

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

func (m *matrixV2) Render() string {
	var builder strings.Builder
	for y := range *m {
		for x := range (*m)[y] {
			builder.WriteString(string((*m)[y][x].color))
			builder.WriteString(string((*m)[y][x].character))
			builder.WriteString(Reset)
		}
		builder.WriteString("\n")
	}
	builder.WriteString(Reset)
	return builder.String()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		// Handle keyboard and mouse events in the viewport
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)

	case frameMsg:

		if !m.ready || m.done {
			cmds = append(cmds, searchTick())
			break
		}

		if !m.done {
			found := startSearch(m.xFocus, m.yFocus, m.grid, 0, &m.searchFuncs)
			m.total += found

			if found > 0 {
				(*m.grid)[m.yFocus][m.xFocus].color = Red
			}

			if m.xFocus == len((*m.grid)[m.xFocus])-1 && m.yFocus == len((*m.grid))-1 {
				m.done = true // We are at the end of the grid
			} else if m.xFocus < len((*m.grid)[m.xFocus])-1 {
				m.xFocus++ // Stay on same line
			} else {
				// Move to next line
				m.xFocus = 0
				m.yFocus++
			}
			m.viewport.SetContent(m.grid.Render())

			cmds = append(cmds, searchTick())
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.grid.Render())
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

func (m *model) View() (s string) {
	if !m.ready {
		return "\n  Initializing..."
	}

	s += fmt.Sprintf("%s\n", m.headerView())
	s += fmt.Sprintf("%s\n", m.viewport.View())
	s += m.footerView()

	return s
}

type frameMsg struct{}

func searchTick() tea.Cmd {
	return tea.Tick(time.Second/280, func(_ time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m *model) Init() tea.Cmd {
	return searchTick()
}

func (m model) headerView() string {
	title := titleStyle.Render(fmt.Sprintf("Finding XMAS! Current: %d", m.total))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
