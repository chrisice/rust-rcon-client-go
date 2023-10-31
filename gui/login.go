// gui/login.go
package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type LoginForm struct {
	ServerURLInput    *widget.Entry
	RconPasswordInput *widget.Entry
	LoginButton       *widget.Button
	StatusLabel       *widget.Label
}

func NewLoginForm() *fyne.Container {
	loginForm := &LoginForm{}

	loginForm.ServerURLInput = widget.NewEntry()
	loginForm.ServerURLInput.SetPlaceHolder("Enter Server URL")

	loginForm.RconPasswordInput = widget.NewPasswordEntry()
	loginForm.RconPasswordInput.SetPlaceHolder("Enter RCON Password")

	loginForm.LoginButton = widget.NewButton("Login", func() {
		// Handle login logic here
		// For now, let's just update the status label
		// In a real implementation, you'd use the entered details to try and establish a connection
		loginForm.StatusLabel.SetText("Attempting to login...")
	})

	loginForm.StatusLabel = widget.NewLabel("")

	return container.NewVBox(
		loginForm.ServerURLInput,
		loginForm.RconPasswordInput,
		loginForm.LoginButton,
		loginForm.StatusLabel,
	)
}
