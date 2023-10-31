package models

// User represents a player or user connected to the Rust server.
type User struct {
	ID       string // Maybe the SteamID or any unique identifier.
	Username string
	IP       string
	Status   string // E.g., "Online", "Offline", "Playing", etc.
	// Add other fields as needed.
}

// NewUser creates a new User instance.
func NewUser(id string, username string, ip string, status string) *User {
	return &User{
		ID:       id,
		Username: username,
		IP:       ip,
		Status:   status,
	}
}
