package database

// Database provides an interface to store and retrieve shortened URLs
type Database interface {
	Add(string, string) error
	Get(string) (string, error)
}
