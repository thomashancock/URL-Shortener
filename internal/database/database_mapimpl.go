package database

import (
	"errors"
)

// database_mapimpl implements the Database interface using a map
type database_mapimpl struct {
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
