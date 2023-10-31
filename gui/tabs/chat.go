// gui/tabs/chat.go
package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ChatTab represents the GUI elements for the chat tab.
type ChatTab struct {
	ChatDisplay *widget.Entry  // This will be used to display chat messages.
	InputField  *widget.Entry  // Input field for typing a message.
	SendButton  *widget.Button // Button to send the message.
}

// NewChatTab creates and returns a new chat tab.
func NewChatTab() fyne.CanvasObject {
	// Create the components for the chat tab.
	chat := &ChatTab{}

	chat.ChatDisplay = widget.NewMultiLineEntry()
	chat.ChatDisplay.Disable()
	chat.InputField = widget.NewEntry()
	chat.SendButton = widget.NewButton("Send", chat.sendMessage)

	// Arrange them in a VBox (vertical box layout).
	content := container.NewVBox(
		chat.ChatDisplay,
		chat.InputField,
		chat.SendButton,
	)

	return content
}

// sendMessage is the function that will handle sending a chat message.
func (c *ChatTab) sendMessage() {
	// Here, you'd integrate with your RCON client to send the message.
	// For this example, we're just appending it to the chat display.
	message := c.InputField.Text
	c.ChatDisplay.SetText(c.ChatDisplay.Text + "\n" + message)
	c.InputField.SetText("") // Clear the input field.
}
