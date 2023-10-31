// gui/tabs/userList.go
package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// UserListTab represents the GUI elements for the User List tab.
type UserListTab struct {
	UserItems     []fyne.CanvasObject // A slice of widgets to display users.
	UserList      *fyne.Container     // A box to display users.
	RefreshButton *widget.Button      // Button to refresh the user list.
}

// NewUserListTab creates and returns a new User List tab.
func NewUserListTab() fyne.CanvasObject {
	users := &UserListTab{}

	users.UserList = container.New(layout.NewVBoxLayout())
	users.RefreshButton = widget.NewButton("Refresh", users.fetchUsers)

	// Use a VBox layout to arrange widgets vertically.
	content := container.New(layout.NewVBoxLayout(),
		users.UserList,
		users.RefreshButton,
	)

	return content
}

// fetchUsers retrieves the list of users from the server.
func (u *UserListTab) fetchUsers() {
	// Ideally, this function should interact with the Rust server to get the list of users.
	// For now, we'll simulate fetching user data with some placeholder values.
	placeholderUsers := []string{"User1", "User2", "User3"}

	// Clear the UserList.
	for _, user := range placeholderUsers {
		u.UserItems = append(u.UserItems, widget.NewLabel(user))
	}

	// Update the UserList with the new user items.
	u.UserList = container.New(layout.NewVBoxLayout(), u.UserItems...)

	// In a real-world scenario, you'd replace the placeholder logic
	// with a method to fetch the actual user list from the Rust server and update UserList.
}
