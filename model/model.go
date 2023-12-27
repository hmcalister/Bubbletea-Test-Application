package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// A struct conforming to the Bubbletea Model interface.
//
// This model defines the behavior of a TODO list application.
type ApplicationStruct struct {
	// Hold some state about the application

	newItemTextInput textinput.Model
	choices          []string
	cursor           int
	selected         map[int]interface{}
}

// Create a new Application, initializing the values accordingly
func NewApplicationStruct() ApplicationStruct {
	newItemTextInput := textinput.New()
	newItemTextInput.Placeholder = "Add New Item"

	return ApplicationStruct{
		newItemTextInput: newItemTextInput,
		choices:          []string{"Item 1", "Item 2"},
		cursor:           0,
		selected:         make(map[int]interface{}),
	}
}

// Define the Init() function on our model struct.
//
// This function is called by Bubbletea when the app is started,
// so any additional initialization can occur here.
//
// The return type, tea.Cmd, is used for I/O on initialization.
// TODO: Figure out exactly what this means.
func (app ApplicationStruct) Init() tea.Cmd {
	return textinput.Blink
}

// Define the Update() function on our model struct.
//
// This function is called by Bubbletea when the application is
// interacted with. Here we need to update any application state
// that may change due to the incoming interaction.
//
// We return a new ApplicationStruct, one that has been updated,
// and a tea.Cmd which can handle I/O.
func (app ApplicationStruct) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Check what kind of interaction has occurred
	// Note we use a type switch so we can perform logic
	// specific to each kind of message
	switch msg := msg.(type) {
	// Keypress
	case tea.KeyMsg:
		switch msg.String() {
		// Instructions to quit the program.
		case "ctrl+c":
			return app, tea.Quit

		case "q":
			if !app.newItemTextInput.Focused() {
				return app, tea.Quit
			}

		// Handle arrow keys
		case "up":
			// If we pressed up while focused on the new item prompt, defocus it
			if app.newItemTextInput.Focused() {
				app.newItemTextInput.Blur()
			}
			app.cursor = max(app.cursor-1, 0)
		case "down":
			app.cursor = min(app.cursor+1, len(app.choices))
			if app.cursor == len(app.choices) {
				return app, app.newItemTextInput.Focus()
			}

		case " ":
			if !app.newItemTextInput.Focused() {
				// Check if the highlighted item is already selected or not
				_, ok := app.selected[app.cursor]
				if ok {
					delete(app.selected, app.cursor)
				} else {
					app.selected[app.cursor] = struct{}{}
				}
			}

		case "enter":
			if app.newItemTextInput.Focused() {
				textInputValue := app.newItemTextInput.Value()
				if len(textInputValue) == 0 {
					return app, nil
				}

				// We have got a new item to add!
				app.choices = append(app.choices, textInputValue)
				app.newItemTextInput.Reset()
				app.cursor += 1
			} else {
				// Check if the highlighted item is already selected or not
				_, ok := app.selected[app.cursor]
				if ok {
					delete(app.selected, app.cursor)
				} else {
					app.selected[app.cursor] = struct{}{}
				}
			}
		}
	}

	if app.newItemTextInput.Focused() {
		app.newItemTextInput, cmd = app.newItemTextInput.Update(msg)
	}

	// Still not returning any commands.
	return app, cmd
}

// Define the View function on our model struct.
// This function renders the application state to the terminal.
//
// Interestingly, the return type is `string`; we are literally
// just printing a string out!
func (app ApplicationStruct) View() string {
	var appStringBuilder strings.Builder
	appStringBuilder.WriteString(headerStyle.Render("TODO List Application") + "\n\n")

	// For each choice:
	for index, choice := range app.choices {
		var currentItemStringBuilder strings.Builder
		// If the cursor is on this choice, add it, but default to no cursor
		cursor := "  "
		if index == app.cursor {
			cursor = ">>"
		}

		// If the choice is selected, mark it so, but default to no mark
		checked := " "
		if _, ok := app.selected[index]; ok {
			checked = "x"
		}

		itemStr := fmt.Sprintf("%s [%s] %s", cursor, checked, choice)
		if index == app.cursor {
			currentItemStringBuilder.WriteString(activeListItemStyle.Render(itemStr))
		} else {
			currentItemStringBuilder.WriteString(inactiveListItemStyle.Render(itemStr))
		}
		currentItemString := currentItemStringBuilder.String()

		// Add this item to the appString
		appStringBuilder.WriteString(listItemBorder.Render(currentItemString) + "\n")
	}

	// Add the text input to the bottom of the app
	appStringBuilder.WriteString(textInputStyle.Render(app.newItemTextInput.View()))

	// // Finally, add a footer
	// if app.newItemTextInput.Focused() {
	// 	appStringBuilder.WriteString("\n\n")
	// } else {
	// 	appStringBuilder.WriteString("\nPress q to quit.\n")
	// }
	return appStringBuilder.String()
}
