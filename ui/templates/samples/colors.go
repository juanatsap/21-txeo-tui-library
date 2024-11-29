package samples

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	txui "txeo-tui-library/ui"
)

func getColorsSampleHoursUI() string {

	var sb strings.Builder
	// Lista de horas trabajadas
	hoursWorked := []float64{9.0, 8.5, 8.0, 7.5, 7.0, 6.5, 6.0, 5.5, 5.0, 4.5, 4.0, 3.5, 3.0, 2.5, 2.0, 1.5, 1.0, 0.5}

	for _, hours := range hoursWorked {
		hoursFormatted := fmt.Sprintf("Horas trabajadas: %.1f", hours)
		colorForHours := txui.GetBackgroundColorForHours(hours)
		if hours > 8.0 {
			colorForHours = "#ff0000"
		}
		if hours < 0.5 {
			txui.HoursDistributionStyle = txui.HoursDistributionStyle.Foreground(lipgloss.Color("#000000"))
		}
		sb.WriteString("\n" + txui.HoursDistributionStyle.Background(lipgloss.Color(colorForHours)).Render(hoursFormatted))
	}

	return sb.String()
}
