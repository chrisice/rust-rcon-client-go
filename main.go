package main

import (
	"rust-rcon-client/gui/tabs" // Adjust this import to the correct path of your 'tabs' package

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Rust RCON Client")

	tabs := container.NewAppTabs(
		container.NewTabItem("Console", tabs.NewConsoleTab()),
		container.NewTabItem("Chat", tabs.NewChatTab()),
		container.NewTabItem("Oxide Logs", tabs.NewOxideLogsTab()),
		container.NewTabItem("User List", tabs.NewUserListTab()),
		container.NewTabItem("Settings", tabs.NewSettingsTab()),
	)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(800, 600)) // Set width to 800 and height to 600
	myWindow.ShowAndRun()

}
