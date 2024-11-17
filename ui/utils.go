package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

func ResetColor() string {
	return "\033[0m"
}
func TruncateString(input string, length int, ellipsis string, finalString string) string {

	if len(input) > length {
		return input[:length-len(ellipsis)] + ellipsis
	}

	var sb strings.Builder
	// Add the rest of the spaces and  then the final string
	for i := len(input); i < length; i++ {
		sb.WriteString(" ")
	}

	return fmt.Sprint(input+sb.String(), string(Reset), finalString)
}
func WaitForLoading() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(1 * time.Second)
		return "done"
	}
}
func GetLanguageCode(language string) string {
	switch language {
	case "Inglés":
		return "en"
	case "Español":
		return "es"
	case "Portugues":
		return "pt"
	case "Francés":
		return "fr"
	case "Alemán":
		return "de"
	case "Italiano":
		return "it"
	case "Chino":
		return "zh"
	case "Japones":
		return "ja"
	}
	return "en"
}
func GetLanguageFlag(language string) string {
	switch language {
	case "Inglés":
		return "🇬🇧"
	case "Español":
		return "🇪🇸"
	case "Portugues":
		return "🇧🇷"
	case "Francés":
		return "🇫🇷"
	case "Alemán":
		return "🇩🇪"
	case "Italiano":
		return "🇮🇹"
	case "Chino":
		return "🇨🇳"
	case "Japones":
		return "🇯🇵"
	}
	return "🇬🇧"
}

// Generate a blend of colors.
func makeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))))
	}
	return
}

// Convert a colorful.Color to a hexadecimal format.
func colorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}

// Helper function for converting colors to hex. Assumes a value between 0 and
// 1.
func colorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}
