package database

import (
	"testing"
	"os"

	"../utils"

	"github.com/stretchr/testify/assert"
)

func clean(t *testing.T, file string) {
	if fileExists(file) {
		err := os.Remove(file)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// Test_Database tests the basic functionality of the database
func Test_Database(t *testing.T) {
	testLog := &utils.TestLogger{}
	dbFile := "sqlite-test.db"

	// Remove file if it already exists
	clean(t, dbFile)

	testDB := NewSQLDatabase(testLog, dbFile)
	expectedLog := "Created SQL DB\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)

	nEntries, err := testDB.NEntries()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nEntries, 0)

	// Add entry to DB
	testDB.Add("0", "Test")

	nEntries, err = testDB.NEntries()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nEntries, 1)

	// Retreive entry from DB
	entry, err := testDB.Get("0")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Test"
	assert.Equal(t, entry, expected)

	clean(t, dbFile)
}

// Test_Database_NoEntry tests when get is called with no matching entry
func Test_Database_NoEntry(t *testing.T) {
	testLog := &utils.TestLogger{}
	dbFile := "sqlite-test.db"

	// Remove file if it already exists
	clean(t, dbFile)

	testDB := NewSQLDatabase(testLog, dbFile)
	expectedLog := "Created SQL DB\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)

	nEntries, err := testDB.NEntries()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nEntries, 0)

	// Retrieve entry from empty DB
	entry, err := testDB.Get("0")

	assert.NotNil(t, err)
	expectedErr := "Unable to find entry for 0 in db: sql: no rows in result set"
	assert.Equal(t, err.Error(), expectedErr)

	expected := ""
	assert.Equal(t, entry, expected)

	expectedLog = "Unable to find entry for 0 in db\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)

	clean(t, dbFile)
}

// Test_Database_Persistence tests the database can be loaded correctly
func Test_Database_Persistence(t *testing.T) {
	testLog := &utils.TestLogger{}
	dbFile := "sqlite-test.db"

	// Remove file if it already exists
	clean(t, dbFile)

	testDB := NewSQLDatabase(testLog, dbFile)
	expectedLog := "Created SQL DB\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)

	// Check database is empty
	nEntries, err := testDB.NEntries()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, nEntries, 0)

	// Add an entry
	testDB.Add("0", "Test")

	// Open database again
	testDB_New := NewSQLDatabase(testLog, dbFile)

	// Retreive persisted entry
	entry, err := testDB_New.Get("0")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Test"
	assert.Equal(t, entry, expected)

	clean(t, dbFile)
}
