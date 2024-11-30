package ui

import "github.com/charmbracelet/bubbles/textinput"

func InitTI() textinput.Model {
	// Crear el campo de texto para la introducción manual de tablero y mes
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 50
	ti.Prompt = ""
	return ti
}
