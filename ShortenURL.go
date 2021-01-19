package main

import (
	"strconv"
)


func GetShortURL(val int) string {
	// For now, just turn the integer into a string
	out := strconv.Itoa(val)

	return out
}
