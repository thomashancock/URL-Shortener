package database

import (
	"errors"
)

type Database interface {
	Add(string, string) error
	Get(string) (string, error)
}

type database struct {
	db map[string]string
}

func (db *database) Add(path string, redirect string) (error) {
	if db.db[path] != "" {
		return errors.New("Path already in database")
	}
	db.db[path] = redirect
	return nil
}

func (db *database) Get(path string) (string, error) {
	redirect := db.db[path]
	if redirect == "" {
		return "", errors.New("Path not found")
	}
	return redirect, nil
}

func NewDatabase() Database {
	return &database{db: make(map[string]string)}
}
