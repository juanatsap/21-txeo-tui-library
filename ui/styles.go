package ui

import (
	"fmt"
	"hash/fnv"
	"image/color"
	"net/http"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/logrusorgru/aurora"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Color map for hours to colors
var ColorMap = map[float64]string{
	8.5: "#FF0000", // Pure Red (maximum danger)
	8.0: "#FF0000", // Pure Red
	7.5: "#D9534F", // Darker red to create more distinction from pure red
	7.0: "#EC6B4E",
	6.5: "#F07848",
	6.0: "#F28445",
	5.5: "#F49142",
	5.0: "#F6964C",
	4.5: "#F7A352",
	4.0: "#F8AD58", // Transition to orange
	3.5: "#F8B665",
	3.0: "#F8BA6A",
	2.5: "#F8C077",
	2.0: "#F8C680",
	1.5: "#F8CB88",
	1.0: "#F8D091",
	0.5: "#F8D495",
	0.0: "", // Pastel Orange (safety)
}

// Color map for latencies in seconds (soft pastel colors)
var LatencyColorMap = map[float64]string{
	0.0:  "#FFFFFF", // Pure White (ideal)
	0.25: "#FFF9E6", // Very Light Yellow
	0.5:  "#FFF3CC", // Soft Yellow
	0.75: "#FFECCC", // Pale Yellow
	1.0:  "#FFE6B3", // Light Yellow
	1.25: "#FFDF99", // Soft Pastel Yellow
	1.5:  "#FFD980", // Medium Pastel Yellow
	1.75: "#FFD366", // Yellow-Orange
	2.0:  "#FFCC66", // Light Orange
	2.25: "#FFB84D", // Pastel Orange
	2.5:  "#FFA333", // Softer Orange-Red
	2.75: "#FF8F33", // Soft Red-Orange
	3.0:  "#FF7A33", // Muted Orange-Red
	3.25: "#FF664D", // Soft Red
	3.5:  "#FF5233", // Medium Red
	3.75: "#FF3D33", // Muted Red
	4.0:  "#FF2922", // Deeper Red
	4.25: "#FF1500", // Deep Red
	4.5:  "#E60000", // Strong Softened Red
	4.75: "#CC0000", // Very Deep Red
	5.0:  "#990000", // Pure Dark Red (Critical)
}

var TitleCaser = cases.Title(language.Und)

// ANSI Color definitions.
var (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Black   = "\033[30m"
	Gray    = "\033[37m"

	DarkYellow  = "\033[33;43m"
	DarkBlue    = "\033[34;44m"
	DarkRed     = "\033[31;41m"
	DarkGreen   = "\033[32;42m"
	DarkCyan    = "\033[36;46m"
	DarkMagenta = "\033[35;45m"
	DarkWhite   = "\033[37;47m"

	// ANSI escape codes for formatting
	Reset           = "\033[0m" // Reset to default color
	Bold            = "\033[1m"
	BoldRed         = "\033[1;31m"
	BoldBlue        = "\033[1;34m"
	BoldCyan        = "\033[1;37m"
	BoldOrange      = "\033[1;33m"
	BoldGreen       = "\033[1;32m"
	BoldYellow      = "\033[1;33m"
	BoldMagenta     = "\033[1;35m"
	Underline       = "\033[4m"
	UnderlineRed    = "\033[4;31m"
	UnderlineBlue   = "\033[4;34m"
	UnderlineOrange = "\033[4;33m"
	UnderlineGreen  = "\033[4;32m"
	UnderlineYellow = "\033[4;33m"
	BoldDarkYellow  = "\033[1;33;43m"
)

// HEX Color definitions.
var (
	HEX = map[string]string{
		"Green":   "#71fd4b",
		"Red":     "#fd4b4b",
		"Yellow":  "#fdcc4b",
		"Blue":    "#4b4bfd",
		"Magenta": "#fd4bfd",
		"Cyan:":   "#4bffd9",
		"White":   "#ffffff",
		"Black":   "#000000",
		"Gray":    "#383838",

		"DarkYellow":  "#fdcc4b",
		"DarkBlue":    "#4b4bfd",
		"DarkRed":     "#fd4b4b",
		"DarkGreen":   "#71fd4b",
		"DarkCyan":    "#74ade9",
		"DarkMagenta": "#fd4bfd",
		"DarkWhite":   "#ffffff",

		"LightMagenta": "#f9d5ff",

		"TxeoCalculatorGreen": "#43BF6D",
		"GooglePlusRed":       "#c27a71",
		"GoogleBlue":          "#78a4ea",
		"AndroidGreen":        "#93d396",
	}
)

// Style definitions.
var (
	// General.
	PanelBackgroundColor = lipgloss.AdaptiveColor{Light: "#71fd4b", Dark: "#031935ff"}
	Subtle               = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	Highlight            = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	Special              = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	ErrorColor           = lipgloss.AdaptiveColor{Light: "#BF616A", Dark: "#F07178"}
	Divider              = lipgloss.NewStyle().SetString("•").Padding(0, 1).Foreground(Subtle).String()
	Url                  = lipgloss.NewStyle().Foreground(Special).Render

	// Page.
	DocStyle = lipgloss.NewStyle().Padding(1, 1).Margin(0).Align(lipgloss.Center)

	// Tabs.
	ActiveTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
	TabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
	PanelBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	RegularTab        = lipgloss.NewStyle().Border(TabBorder, true).BorderForeground(Highlight).Padding(0, 1)
	RegularPanelStile = RegularTab
	ActiveTab         = RegularTab.Border(ActiveTabBorder, true).Bold(true)
	TabGap            = RegularTab.BorderTop(false).BorderLeft(false).BorderRight(false)

	// Paragraphs/History.
	PanelStyle = lipgloss.NewStyle().Border(PanelBorder).Align(lipgloss.Center).Foreground(lipgloss.Color("#FAFAFA")).Margin(0, 1).Padding(1, 2)

	// Crear un estilo base
	ActiveStyle    = lipgloss.NewStyle().BorderForeground(lipgloss.Color("120")).Align(lipgloss.Center).BorderStyle(lipgloss.DoubleBorder())
	NonActiveStyle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("240")).Align(lipgloss.Center).UnsetBorderStyle()
	// NonBorderedActiveStyle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("240")).BorderStyle(lipgloss.HiddenBorder())

	// Title.
	TitleStyle = lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#43BF6D")).Bold(true).Margin(2, 0, 0, 0)
	DescStyle  = lipgloss.NewStyle().Align(lipgloss.Left).MarginTop(1).Foreground(lipgloss.Color("#874BFD")).Inline(true)
	InfoStyle  = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderTop(true).BorderForeground(Subtle)

	// Dialog.
	DialogBoxStyle    = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#874BFD")).Padding(1, 0).BorderTop(true).BorderLeft(true).BorderRight(true).BorderBottom(true)
	ButtonStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF7DB")).Background(lipgloss.Color("#888B7E")).Padding(0, 3).MarginTop(1)
	ActiveButtonStyle = ButtonStyle.Foreground(lipgloss.Color("#FFF7DB")).Background(lipgloss.Color("#e581a6")).MarginRight(2).Underline(true)
	WidthInfoStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#90747e")).Bold(true).Foreground(lipgloss.Color("#ffffff"))

	// List.
	ListExample = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, true, false, false).BorderForeground(Subtle).MarginRight(2).Height(8)
	ListHeader  = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).BorderForeground(Subtle).MarginRight(2).Render
	ListItem    = lipgloss.NewStyle().PaddingLeft(2).Render
	CheckMark   = lipgloss.NewStyle().SetString("✓").Foreground(Special).PaddingRight(1).String()
	ListDone    = func(s string) string {
		return CheckMark + lipgloss.NewStyle().Strikethrough(true).Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).Render(s)
	}

	// Status Bar.
	StatusNugget                 = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1)
	StatusBarStyle               = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})
	StatusStyle                  = lipgloss.NewStyle().Inherit(StatusBarStyle).Foreground(lipgloss.Color("#FFFDF5")).Background(lipgloss.Color("#3f8edd")).Bold(true).Padding(0, 1).MarginRight(1)
	CurrentBedFilesBoardStyle    = StatusNugget.Foreground(lipgloss.AdaptiveColor{Dark: "#343433", Light: "#C1C6B2"}).Background(lipgloss.Color("#e4f0f8")).Align(lipgloss.Right).Bold(true)
	StatusText                   = lipgloss.NewStyle().Inherit(StatusBarStyle)
	CurrentBedFilesListStyle     = StatusNugget.Background(lipgloss.Color("#377ec4"))
	CurrentBedFilesUsernameStyle = StatusNugget.Background(lipgloss.Color("#4988c8ae"))

	// Hours Distribution.
	HoursDistributionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#b36969")).
				Margin(0, 1).
				Padding(0, 5).Align(lipgloss.Center)

	CalendarHoursDistributionStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FFFFFF")).
					Background(lipgloss.Color("#b36969"))

	// General styles
	// Colores inspirados en BedFiles
	BedFilesBlue       = "#0079BF"
	BedFilesBackground = "#43BF6D"

	// Estilo principal para el contenido
	Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color(BedFilesBlue)).
		Background(lipgloss.Color(BedFilesBackground)).
		Padding(1, 2).
		Margin(1).
		Align(lipgloss.Center)

	// Estilo para el texto del usuario y sugerencias
	UserInputStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))  // Blanco brillante
	SuggestionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244")) // Gris claro

	// Bordered styles
	BaseBorderedStyle    = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).Align(lipgloss.Left).Padding(1, 4)
	CompactBorderedStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#c8c8c8")).Align(lipgloss.Left).Padding(0, 4)

	// Dialog and Base styles
	BaseStyle       = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	NoBorderedStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderTop(false).BorderLeft(false).BorderRight(false).BorderBottom(false)
	DialogStyle     = lipgloss.NewStyle().Width(50).Align(lipgloss.Center)
)

func GetStringInColor(color string, s string) string {
	colorHex := HEX[color]
	return lipgloss.NewStyle().Foreground(lipgloss.Color(colorHex)).Render(s)
}

/* ╭──────────────────────────────────────────╮ */
/* │         BUBBLE TABLE DEFINITION          │ */
/* ╰──────────────────────────────────────────╯ */
var (
	StyleSubtle = lipgloss.NewStyle().Foreground(lipgloss.Color("#888"))

	StyleBase = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a7a")).
			BorderForeground(lipgloss.Color("#a38")).
			Align(lipgloss.Left).Padding(0, 1)

	StyleBaseRow = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a7a")).
			BorderForeground(lipgloss.Color("#a38")).
			Align(lipgloss.Left).Padding(0, 1)
	StyleCentered = lipgloss.NewStyle().Align(lipgloss.Center).Padding(0, 1)
	StyleLeft     = lipgloss.NewStyle().Align(lipgloss.Left).Margin(0, 1)
	StyleRight    = lipgloss.NewStyle().Align(lipgloss.Right).Padding(0, 1)
)

const (
	ColorNormal   = "#fa0"
	ColorFire     = "#f64"
	ColorElectric = "#ff0"
	ColorWater    = "#44f"
	ColorPlant    = "#8b8"
)

/* ╭──────────────────────────────────────────╮ */
/* │         BUBBLE VIEWS DEFINITIONS         │ */
/* ╰──────────────────────────────────────────╯ */
const (
	ProgressBarWidth  = 71
	ProgressFullChar  = "█"
	ProgressEmptyChar = "░"
	DotChar           = " • "
)

// General stuff for styling the view
var (
	KeywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	SubtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	TicksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	ProgressEmpty = SubtleStyle.Render(ProgressEmptyChar)
	DotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(DotChar)
	MainStyle     = lipgloss.NewStyle().MarginLeft(2)

	// Gradient colors we'll use for the progress bar
	Ramp = makeRampStyles("#B14FFF", "#00FFA3", ProgressBarWidth)
)

/* ╭──────────────────────────────────────────╮ */
/* │                  TRELLO                  │ */
/* ╰──────────────────────────────────────────╯ */
var (
	// Status Bar.
	CurrentTrelloBoardStyle    = StatusNugget.Foreground(lipgloss.AdaptiveColor{Dark: "#343433", Light: "#C1C6B2"}).Background(lipgloss.Color("#e4f0f8")).Align(lipgloss.Right).Bold(true)
	CurrentTrelloListStyle     = StatusNugget.Background(lipgloss.Color("#377ec4"))
	CurrentTrelloUsernameStyle = StatusNugget.Background(lipgloss.Color("#4988c8ae"))
	// General styles
	// Colores inspirados en Trello
	TrelloBlue       = "#0079BF"
	TrelloBackground = "#EBECF0"

	ViewportTitleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
)

var (
	Term      = termenv.EnvColorProfile()
	SubtleII  = makeFgStyle("241")
	Dot       = ColorFg(" • ", "236")
	HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)

func ColorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}
func ColorFloatToHex(f float64) string {
	return colorFloatToHex(f)
}
func MakeRamp(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(ColorToHex(c))))
	}
	return
}
func ColorFg(val, color string) string {
	return termenv.String(val).Foreground((Term.Color(color))).String()
}
func makeFgStyle(color string) func(...string) string {
	return lipgloss.Style{}.Foreground(lipgloss.Color(color)).Render
}

// hashRGB genera un color (pseudo-aleatorio) en [0..255] para R,G,B
func hashRGB(s string) (r, g, b byte) {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	sum := h.Sum32()
	r = byte(sum >> 16)
	g = byte(sum >> 8)
	b = byte(sum)
	return
}

// brightness calcula una aproximación de la luminancia
func brightness(r, g, b byte) float64 {
	return 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
}

// lighten opcional: si quieres asegurar que no sea muy oscuro
func Lighten(c color.RGBA, factor float64) color.RGBA {
	// factor: 0.0 (sin cambio), 0.5 (ilumina 50%), etc.
	rr := float64(c.R)
	gg := float64(c.G)
	bb := float64(c.B)

	rr = rr + (255-rr)*factor
	gg = gg + (255-gg)*factor
	bb = bb + (255-bb)*factor

	return color.RGBA{
		R: uint8(rr),
		G: uint8(gg),
		B: uint8(bb),
		A: 255,
	}
}

// coloredString usa Lipgloss para convertir el IP en un 'badge' coloreado
func ColoredString(stringToColor string) string {
	r, g, b := hashRGB(stringToColor)
	// Opcional: iluminar para evitar colores muy oscuros
	c := Lighten(color.RGBA{R: r, G: g, B: b, A: 255}, 0.4)
	r, g, b = c.R, c.G, c.B

	// Generamos string “#RRGGBB” para Lipgloss
	bgHex := fmt.Sprintf("#%02X%02X%02X", r, g, b)
	bgColor := lipgloss.Color(bgHex)

	// Determina foreground (blanco/negro) dependiendo de la luminosidad
	var fgHex string
	if brightness(r, g, b) > 128 {
		fgHex = "#000000"
	} else {
		fgHex = "#FFFFFF"
	}

	// Creamos estilo Lipgloss
	style := lipgloss.NewStyle().
		Background(bgColor).
		Foreground(lipgloss.Color(fgHex)).
		Bold(true).
		Padding(0, 1) // espacio extra a los lados

	// Renderizamos el IP
	return style.Render(stringToColor)
}

func InterpolateHexColor(color1, color2 string, t float64) string {
	// Convierte colores hex a RGB
	r1, g1, b1 := hexToRGB(color1)
	r2, g2, b2 := hexToRGB(color2)

	// Interpola cada componente
	r := uint8(float64(r1)*(1-t) + float64(r2)*t)
	g := uint8(float64(g1)*(1-t) + float64(g2)*t)
	b := uint8(float64(b1)*(1-t) + float64(b2)*t)

	// Retorna como hex
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// GenerateGradient creates a gradient of colors between start and end in `steps` steps.
func GenerateGradient(startHex, midHex, endHex string, steps int, lightenFactor float64) []string {
	startR, startG, startB := hexToRGB(startHex)
	midR, midG, midB := hexToRGB(midHex)
	endR, endG, endB := hexToRGB(endHex)

	gradient := make([]string, steps)
	midpoint := steps / 2

	for i := 0; i < steps; i++ {
		var r, g, b uint8
		if i < midpoint {
			// Interpolate between start and mid
			t := float64(i) / float64(midpoint)
			r = uint8(float64(startR)*(1-t) + float64(midR)*t)
			g = uint8(float64(startG)*(1-t) + float64(midG)*t)
			b = uint8(float64(startB)*(1-t) + float64(midB)*t)
		} else {
			// Interpolate between mid and end
			t := float64(i-midpoint) / float64(steps-midpoint)
			r = uint8(float64(midR)*(1-t) + float64(endR)*t)
			g = uint8(float64(midG)*(1-t) + float64(endG)*t)
			b = uint8(float64(midB)*(1-t) + float64(endB)*t)
		}

		// Lighten the color
		lightened := Lighten(color.RGBA{R: r, G: g, B: b, A: 255}, lightenFactor)
		gradient[i] = fmt.Sprintf("#%02X%02X%02X", lightened.R, lightened.G, lightened.B)
	}

	return gradient
}

// Convert hex color to RGB
func hexToRGB(hex string) (uint8, uint8, uint8) {
	var r, g, b uint8
	fmt.Sscanf(hex, "#%02X%02X%02X", &r, &g, &b)
	return r, g, b
}
func GetStatusColor(status int) aurora.Value {
	var statusColor aurora.Value
	statusSpaced := " " + strconv.Itoa(status) + " "
	switch {
	case status >= http.StatusInternalServerError:
		statusColor = aurora.BgRed(aurora.White(statusSpaced))
	case status >= http.StatusBadRequest:
		statusColor = aurora.BgYellow(aurora.Black(statusSpaced))
	case status >= http.StatusMultipleChoices:
		statusColor = aurora.BgCyan(aurora.Black(statusSpaced))
	default:
		statusColor = aurora.BgGreen(aurora.Black(statusSpaced))
	}
	return statusColor
}
