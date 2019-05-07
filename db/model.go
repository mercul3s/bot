package db

// User represents slack user information
type User struct {
	Handle string
	ID     int
	Type   string
}

// Config holds bot configuration data.
