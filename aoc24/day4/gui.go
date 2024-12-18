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

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func (m *matrixV2) Render() (output string) {
	for y := range *m {
		for x := range (*m)[y] {
			output += string((*m)[y][x].color)
			output += string((*m)[y][x].character)
			output += Reset
		}
		output += "\n"
	}
	output += Reset
	return
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		// Handle keyboard and mouse events in the viewport
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)

	case frameMsg:
		cmds = append(cmds, searchTick())

		if !m.ready {
			break
		}
		if m.done {
			break
		}

		found := startSearch(m.xFocus, m.yFocus, m.grid, 0)
		m.total += found

		if found > 0 {
			(*m.grid)[m.yFocus][m.xFocus].color = Red
		}

		if m.xFocus == len((*m.grid)[m.xFocus])-1 && m.yFocus == len((*m.grid))-1 {
			// We are at the end of the grid
			m.done = true
		} else if m.xFocus < len((*m.grid)[m.xFocus])-1 {
			// Stay on same line
			m.xFocus++
		} else {
			// Move to next line
			m.xFocus = 0
			m.yFocus++
		}
		m.viewport.SetContent(m.grid.Render())

		cmds = append(cmds, searchTick())

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.grid.Render())
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	var s string

	// New Stuff
	if !m.ready {
		return "\n  Initializing..."
	}

	s += fmt.Sprintf("%s\n", m.headerView())
	s += fmt.Sprintf("%s\n", m.viewport.View())
	s += fmt.Sprintf("%s", m.footerView())

	return s

	// // The header
	// s := "Lets save XMAS?\n\n"

	// // Iterate over our choices
	// s += m.grid.render()

	// s += fmt.Sprintf("\nTotal Found: %d\n", m.total)

	// // The footer
	// s += "\nPress q to quit.\n"

	// // Send the UI for rendering
	// return s
}

type frameMsg struct{}

func searchTick() tea.Cmd {
	return tea.Tick(time.Second*10, func(_ time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m *model) Init() tea.Cmd {
	return searchTick()
}

func (m model) headerView() string {
	title := titleStyle.Render("Finding XMAS!")
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
