package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/thomashancock/URL-Shortener/internal/core"

	_ "github.com/mattn/go-sqlite3"
)

// database_sqlimpl implements the Database interface using a map
type database_sqlimpl struct {
	log core.Logger
	db *sql.DB
}

// Add adds an entry to the database
func (d *database_sqlimpl) Add(path string, redirect string) error {
	insertSQL := `INSERT INTO aliases(shorturl, redirect) VALUES (?, ?)`
	statement, err := d.db.Prepare(insertSQL)
	if err != nil {
		d.log.Errorln("Unable to insert to SQL db")
		return fmt.Errorf("Unable to insert to SQL db: %w", err)
	}

	_, err = statement.Exec(path, redirect)
	if err != nil {
		d.log.Errorln("Unable to insert to SQL db")
		return fmt.Errorf("Unable to insert to SQL db: %w", err)
	}

	return nil
}

// Get retreives an entry to the database
func (d *database_sqlimpl) Get(path string) (string, error) {
	// Query DB
	row := d.db.QueryRow("SELECT shorturl, redirect FROM aliases WHERE shorturl=?", path)

	// Get information from row
	var shorturl string
	var redirect string
	err := row.Scan(&shorturl, &redirect)
	if err != nil {
		d.log.Infof("Unable to find entry for %s in db\n", path)
		return "", fmt.Errorf("Unable to find entry for %s in db: %w", path, err)
	}

	return redirect, nil
}

// NEntries returns the number of entries in the database
func (d *database_sqlimpl) NEntries() (int, error) {
	d.log.Infoln("Counting entries in DB")

	var count int
	_ = d.db.QueryRow("SELECT COUNT(*) FROM aliases").Scan(&count)
	return count, nil
}

// fileExists checks if a file exists on the system
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// init initialises the database
func (d *database_sqlimpl) init() {
	d.log.Infoln("Creating SQL DB")

	if d.db == nil {
		d.log.Fatalln("Tried to create database on nil pointer")
	}

	createTableSQL := `CREATE TABLE aliases (
		"index" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"shorturl" text UNIQUE,
		"redirect" text
	 	);`
	statement, err := d.db.Prepare(createTableSQL) // Prepare SQL Statement
	if err != nil {
		d.log.Fatalln("Unable to prepare SQL db")
	}
	statement.Exec()
	d.log.Infoln("Created SQL DB")
}

// open opens the DB
// if the file doesn't exist, it is created and initialised
func (d *database_sqlimpl) open(dbFile string) {
	// If file doesn't exist, defer database initalisation
	if !fileExists(dbFile) {
		defer d.init()
	}

	d.log.Infoln("Opening SQL DB")
	var err error
	d.db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		d.log.Fatalf("Error opening db: %v", err)
	}
}

// NewDatabase creates a new database
func NewSQLDatabase(log core.Logger, dbFile string) Database {
	db := database_sqlimpl{log: log, db: nil}
	db.open(dbFile)
	return &db
}
