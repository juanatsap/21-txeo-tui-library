package widgets

import (
	"fmt"
	"math"
	"strings"
	txui "txeo-tui-library/ui"
)

const (
	ProgressBarWidth  = 71
	ProgressFullChar  = "█"
	ProgressEmptyChar = "░"
	DotChar           = " • "
)

var (
	progressEmpty = txui.SubtleII(ProgressEmptyChar)
	ramp          = txui.MakeRamp("#B14FFF", "#00FFA3", ProgressBarWidth)
)

func Checkbox(label string, checked bool) string {
	if checked {
		return txui.ColorFg("[x] "+label, "212") // Color 212 is green (#00FF00)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func ProgressBar(percent float64) string {
	w := float64(ProgressBarWidth)
	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += ramp[i].Render(ProgressFullChar)
	}
	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)
	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}
