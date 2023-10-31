package rcon

// Define constants for each RCON command, which can be handy to avoid mistyping commands
const (
	CommandListPlayers = "listplayers"
	CommandServerInfo  = "serverinfo"
	// ... more commands ...
)

// SendListPlayers sends the listplayers command and retrieves a list of players.
func (c *Client) SendListPlayers() ([]string, error) {
	response, err := c.SendCommand(CommandListPlayers)
	if err != nil {
		return nil, err
	}

	players := parsePlayers(response)
	return players, nil
}

// SendServerInfo sends the serverinfo command and retrieves server-related information.
func (c *Client) SendServerInfo() (ServerInfo, error) {
	response, err := c.SendCommand(CommandServerInfo)
	if err != nil {
		return ServerInfo{}, err
	}

	info := parseServerInfo(response)
	return info, nil
}

// ... more functions for other commands ...

// Utility functions to parse responses from the server
func parsePlayers(response string) []string {
	// Logic to parse the player list from the response string
	return nil // Placeholder, replace with actual logic.
}

func parseServerInfo(response string) ServerInfo {
	// Logic to parse the server info from the response string
	return ServerInfo{} // Placeholder, replace with actual logic.
}

// Define data structures for parsed responses if needed
type ServerInfo struct {
	// Fields representing server information
	// e.g. Name, IP, MaxPlayers, etc.
}
