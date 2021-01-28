package database

import (
	"errors"

	"../core"
)

// database_mapimpl implements the Database interface using a map
type database_mapimpl struct {
	log core.Logger
	db map[string]string
}

// Add adds an entry to the database (map)
func (db *database_mapimpl) Add(path string, redirect string) error {
	if db.db[path] != "" {
		return errors.New("Path already in database")
	}
	db.db[path] = redirect
	return nil
}

// Get retreives an entry to the database (map)
func (db *database_mapimpl) Get(path string) (string, error) {
	redirect := db.db[path]
	if redirect == "" {
		return "", errors.New("Path not found")
	}
	return redirect, nil
}

// NEntries returns the number of entries in the map
func (db *database_mapimpl) NEntries() (int, error) {
	return len(db.db), nil
}

// NewDatabase creates a new database
func NewTestDatabase(log core.Logger) Database {
	return &database_mapimpl{
		log: log,
		db:  make(map[string]string),
	}
}
