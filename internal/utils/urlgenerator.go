package utils

import (
	"strconv"

	"github.com/thomashancock/URL-Shortener/internal/core"
)

// URLGenerator defined an interface for generating unique URLs
type URLGenerator interface {
	Get() (string, error)
}

// genIntIncr implements the URLGenerator interface
type genIntIncr struct {
	log core.Logger
	count *int
}

// Get returns a new URL
// each generated URL should be unique
func (g genIntIncr) Get() (string, error) {
	url := strconv.Itoa(*g.count)
	g.log.Infof("Generated new short URL: %s\n", url)
	*g.count += 1
	return url, nil
}

// NewURLGenerator returns a struct which matches the URLGenerator interface
func NewURLGenerator(log core.Logger, start int) URLGenerator {
	count := start
	return &genIntIncr{
		log: log,
		count: &count,
	}
}
