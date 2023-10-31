// gui/tabs/console.go
package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ConsoleTab represents the GUI elements for the console tab.
type ConsoleTab struct {
	ConsoleOutput *widget.Entry  // This displays server responses and executed commands.
	CommandField  *widget.Entry  // Where the user types in commands.
	ExecuteButton *widget.Button // Button to send the command to the server.
}

// NewConsoleTab creates and returns a new console tab.
func NewConsoleTab() fyne.CanvasObject {
	console := &ConsoleTab{}

	console.ConsoleOutput = widget.NewMultiLineEntry()
	console.ConsoleOutput.SetPlaceHolder("Console output will appear here.")
	console.ConsoleOutput.Disable()
	console.CommandField = widget.NewEntry()
	console.ExecuteButton = widget.NewButton("Execute", console.executeCommand)

	// Adding a padding around elements
	//paddedOutput := container.NewPadded(console.ConsoleOutput)
	paddedField := container.NewPadded(console.CommandField)

	consolebox := container.NewVScroll(console.ConsoleOutput)
	consolebox.SetMinSize(fyne.NewSize(0, 500))

	// Create a horizontal box for command input and execute button with a spacer in between
	commandBox := container.New(layout.NewGridLayout(2),
		paddedField,
		console.ExecuteButton,
	)

	content := container.New(layout.NewVBoxLayout(),
		consolebox,
		commandBox,
	)

	return content

}

// executeCommand sends the command to the server and displays the response.
func (c *ConsoleTab) executeCommand() {
	// This function should interact with the Rust RCON server.
	// For now, let's simulate command execution by appending it to the console.
	command := c.CommandField.Text
	c.ConsoleOutput.SetText(c.ConsoleOutput.Text + "\n> " + command)
	// Here you would also handle getting the server's response and displaying it.

	c.CommandField.SetText("") // Clear the command input field.
}
