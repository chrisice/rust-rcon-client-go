// gui/tabs/settings.go
package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// other imports as needed
)

func NewSettingsTab() fyne.CanvasObject {
	serverInput := widget.NewEntry()
	serverInput.SetPlaceHolder("Enter Server URL")

	passwordInput := widget.NewPasswordEntry()
	passwordInput.SetPlaceHolder("Enter RCON Password")

	settingsTabContent := container.NewVBox(
		serverInput,
		passwordInput,
		widget.NewButton("Save", func() {
			// TODO: Save these settings somewhere (e.g., a file, database, etc.)
		}),
	)

	return settingsTabContent
}
