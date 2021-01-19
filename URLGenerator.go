package main

import (
	"strconv"
)


// Only expose URLGenerator() so internal logic can be easily changed
func URLGenerator() func()string {
	// URL generation logic can be changed here
	return intIncr()
}


// Each new URL is just the next integer
func intIncr() func()string {
	count := -1
	fn := func() string {
		count += 1
		return strconv.Itoa(count)
	}
	return fn
}
