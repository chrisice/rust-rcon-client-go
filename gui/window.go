// gui/window.go
package gui

import (
	"rust-rcon-client/gui/tabs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type MainWindow struct {
	App  fyne.App
	Win  fyne.Window
	Tabs *container.AppTabs
}

// NewMainWindow initializes and returns the main application window.
func NewMainWindow() *MainWindow {
	application := app.New()
	window := application.NewWindow("Rust RCON Client")

	mainWindow := &MainWindow{
		App: application,
		Win: window,
	}

	// Create tabs (using the functions from other files).
	consoleTab := container.NewTabItem("Console", tabs.NewConsoleTab())
	chatTab := container.NewTabItem("Chat", tabs.NewChatTab())
	oxideLogsTab := container.NewTabItem("Oxide Logs", tabs.NewOxideLogsTab())
	userListTab := container.NewTabItem("User List", tabs.NewUserListTab())
	settingsTab := container.NewTabItem("Settings", tabs.NewSettingsTab())
	// ... Add other tabs as needed

	// Set up tabs.
	mainWindow.Tabs = container.NewAppTabs(consoleTab, chatTab, oxideLogsTab, userListTab, settingsTab)

	// Set window content.
	window.SetContent(mainWindow.Tabs)

	return mainWindow
}

// Run starts the main application loop.
func (w *MainWindow) Run() {
	w.Win.ShowAndRun()
}
