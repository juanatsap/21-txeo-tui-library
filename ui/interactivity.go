package ui

import (
	"fmt"
	"math"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	TickMsg  struct{}
	FrameMsg struct{}
)

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}

func Frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return FrameMsg{}
	})
}

func Checkbox(label string, checked bool) string {
	if checked {
		return CheckboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func Progressbar(percent float64) string {
	w := float64(ProgressBarWidth)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += Ramp[i].Render(ProgressFullChar)
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(ProgressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}
