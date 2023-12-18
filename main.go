package main

import (
	"fmt"
	"hmcalister/bubbleteaTestApp/model"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create a new instance of our ApplicationStruct, implementing the tea.Model interface
	app := model.NewApplicationStruct()

	// Create a Bubbletea program using the new app
	program := tea.NewProgram(app)

	// Run the program, checking for errors as we do
	if _, err := program.Run(); err != nil {
		fmt.Printf("error during bubbletea program.run(): %v", err)
		os.Exit(1)
	}
}
