package database

// Database provides an interface to store and retrieve shortened URLs
type Database interface {
	Add(string, string) error
	Get(string) (string, error)
}

// NewDatabase creates a new database
func NewDatabase() Database {
	return &database_mapimpl{db: make(map[string]string)}

	// var db database_sqlimpl
	// db.init()
	// return &db
}
