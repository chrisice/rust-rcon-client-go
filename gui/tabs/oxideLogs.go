// gui/tabs/oxideLogs.go
package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// OxideLogsTab represents the GUI elements for the Oxide Logs tab.
type OxideLogsTab struct {
	LogDisplay    *widget.Entry  // This displays the Oxide log entries.
	RefreshButton *widget.Button // Button to fetch the latest logs.
}

// NewOxideLogsTab creates and returns a new Oxide Logs tab.
func NewOxideLogsTab() fyne.CanvasObject {
	logs := &OxideLogsTab{}

	logs.LogDisplay = widget.NewMultiLineEntry()
	logs.LogDisplay.Disable()
	logs.RefreshButton = widget.NewButton("Refresh", logs.fetchLogs)

	// Arrange them in a VBox (vertical box layout).
	content := container.NewVBox(
		logs.LogDisplay,
		logs.RefreshButton,
	)

	return content
}

// fetchLogs retrieves the latest Oxide logs from the server.
func (o *OxideLogsTab) fetchLogs() {
	// Ideally, this function should interact with the Rust server to get the Oxide logs.
	// For now, let's simulate fetching logs with a placeholder.
	placeholderLogs := `
[INFO] Plugin loaded.
[ERROR] Something went wrong.
[INFO] Another log entry.
	`
	o.LogDisplay.SetText(placeholderLogs)

	// In a real-world scenario, you'd replace the placeholder with logic
	// to fetch the actual logs from the Rust server and update the LogDisplay.
}
