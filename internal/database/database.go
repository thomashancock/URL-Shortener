package database

import (
	"errors"
)

// Database provides an interface to store and retrieve shortened URLs
type Database interface {
	Add(string, string) error
	Get(string) (string, error)
}

// database implements the Database interface using a map
type database struct {
	db map[string]string
}

// Add adds an entry to the database (map)
func (db *database) Add(path string, redirect string) error {
	if db.db[path] != "" {
		return errors.New("Path already in database")
	}
	db.db[path] = redirect
	return nil
}

// Get retreives an entry to the database (map)
func (db *database) Get(path string) (string, error) {
	redirect := db.db[path]
	if redirect == "" {
		return "", errors.New("Path not found")
	}
	return redirect, nil
}

// NewDatabase creates a new database
func NewDatabase() Database {
	return &database{db: make(map[string]string)}
}
