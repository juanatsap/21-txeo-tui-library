package ui

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/ozgio/strutil"
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

func printEntryMessage() string {
	return "Bienvenido al Dashboard de BedFiles Audios Administrator"
}
func PrintExitMessage() string {
	return "Gracias por usar BedFiles Audios Administrator"
}
func CenterBlockText(text string, width int) string {

	// If width is even, add 1 to the width
	if width%2 == 0 {
		width++
	}
	textCentered := strutil.AlignCenter(text, width)
	return textCentered
}
func ClearScreen() {
	fmt.Print("\033[2J\033[1;1H")
}
func Float64FromString(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
func FormatStringAsFloatWithDecimals(s string, decimals int) string {
	decimalsExpression := fmt.Sprintf("%%.%df", decimals)
	f, _ := strconv.ParseFloat(s, 64)
	return fmt.Sprintf(decimalsExpression, f)
}
func GetBackgroundColorForHours(hours float64) string {
	colorCode, exists := ColorMap[hours]
	if !exists {
		colorCode = "" // Default color (grey) if hours do not match
	}

	// If hours > 8, set the color to red
	if hours > 8 {
		colorCode = "#ff0000"
	}

	if hours < 0 {
		colorCode = "#a2079a"
	}
	return colorCode
}
func SlugifyDate(date string) string {

	fmt.Println("date: ", date)
	// Convert the date string to a time.Time object
	t, _ := time.Parse("2006-01-02", date)

	// Take the year, month, and day
	year := t.Year()
	month := t.Month()
	day := t.Day()

	// Slugify the date
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
func SlugifyDay(year, month, day int) string {
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
func SortMapByValue(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}
func StringToMonth(monthName string) time.Month {
	monthName = strings.ToLower(monthName) // Convert to lowercase for case-insensitive comparison
	switch monthName {
	case "january":
		return time.January
	case "february":
		return time.February
	case "march":
		return time.March
	case "april":
		return time.April
	case "may":
		return time.May
	case "june":
		return time.June
	case "july":
		return time.July
	case "august":
		return time.August
	case "september":
		return time.September
	case "october":
		return time.October
	case "november":
		return time.November
	case "december":
		return time.December
	default:

		// Exit the program if the month name is not recognized
		fmt.Println()
		fmt.Println(BoldRed, "Error:"+Reset+Bold+" Unknown month name", Reset)
		fmt.Println()
		fmt.Println(Bold, "Usage: "+Reset, "trello-calculator <board-name> <month>")
		fmt.Println(Bold, "Example: "+Reset, "trello-calculator livgolf august")
		fmt.Println()
		fmt.Println(Bold, "Valid months: "+Reset, "January, February, March, April, May, June, July, August, September, October, November, December")
		fmt.Println()

		// Exit the program
		os.Exit(1)

		return time.January
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func printExitMessage() string {
	return "Gracias por usar Trello Calculator"
}
func resetColor() string {
	return "\033[0m"
}

/* ╭──────────────────────────────────────────╮ */
/* │             TXEO-CALCULATOR              │ */
/* ╰──────────────────────────────────────────╯ */
// "github.com/alexeyco/simpletable"
// Level 4: Aux
func GetBoxWithShadowEffectUI(icon, text string) string {
	var sb strings.Builder

	if icon == "" {
		icon = "📦"
	}
	icon = icon + " "

	maxLength := 11 + len(text)
	textFixed := fmt.Sprintf("%-"+fmt.Sprint(len(text))+"s", text)

	// sb.WriteString("\n")
	sb.WriteString(" ╭" + strings.Repeat("─", maxLength) + "╮\n")
	sb.WriteString("  │   " + icon + textFixed + "     │ ╮\n")
	sb.WriteString("  ╰" + strings.Repeat("─", maxLength) + "╯ │\n")
	sb.WriteString("    ╰" + strings.Repeat("─", maxLength) + "╯\n")

	boxWithShadowEffectContentStyle := lipgloss.NewStyle().Align(lipgloss.Center)

	return boxWithShadowEffectContentStyle.Render(sb.String())
}

// Pvt functions
func CreateMonthMap(year int, monthName string) map[string]float64 {
	// Get the number of days in the month
	month := StringToMonth(monthName)
	// firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	daysInMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Hour * 24).Day()

	// Create a map to store the days
	monthMap := make(map[string]float64)

	// Populate the map with slugified days
	for day := 1; day <= daysInMonth; day++ {
		slug := SlugifyDay(year, int(month), day)
		monthMap[slug] = 0.0 // Assign 0.0 as the value for each day
	}

	return monthMap
}
func GetTotalHoursThisMonth(hours map[string]float64) float64 {
	totalHours := 0.0
	for _, v := range hours {
		totalHours += v
	}
	return totalHours
}

func IsValidMonth(input string) bool {
	months := []string{"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December"}
	inputLower := strings.ToLower(input)
	for _, month := range months {
		if strings.ToLower(month) == inputLower {
			return true
		}
	}
	return false
}

// Public functions
func AskForMonth() (month string) {
	fmt.Print("Enter month: ")
	fmt.Scanln(&month)
	return month
}
