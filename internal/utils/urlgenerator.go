package utils

import (
	"strconv"
)

// NewURLGenerator returns a function for generating URLs
// Only expose NewURLGenerator() so internal logic can be easily changed
func NewURLGenerator() func() (string, error) {
	// URL generation logic can be changed here
	return intIncr()
}

// Each new URL is just the next integer
func intIncr() func() (string, error) {
	count := -1
	fn := func() (string, error) {
		count += 1
		return strconv.Itoa(count), nil
	}
	return fn
}
