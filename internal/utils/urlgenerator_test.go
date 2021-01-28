package utils

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// Test_URLGenerator tests URLGenerator generates the first expected URL
func Test_URLGenerator(t *testing.T) {
	testLog := &TestLogger{}

	generator := NewURLGenerator(testLog, 0)

	url, err := generator.Get()
	if err != nil {
		t.Fatal(err)
	}

	expected := "0"
	assert.Equal(t, url, expected)

	expectedLog := "Generated new short URL: 0\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}

// Test_URLGenerator_Start tests URLGenerator begins at the right point
func Test_URLGenerator_Start(t *testing.T) {
	testLog := &TestLogger{}

	generator := NewURLGenerator(testLog, 7)

	url, err := generator.Get()
	if err != nil {
		t.Fatal(err)
	}

	expected := "7"
	assert.Equal(t, url, expected)

	expectedLog := "Generated new short URL: 7\n"
	assert.Equal(t, testLog.PrevLog, expectedLog)
}

// Test_URLGenerator tests URLGenerator generates the expected sequence
func Test_URLGenerator_Sequence(t *testing.T) {
	testLog := &TestLogger{}

	num := 4
	generator := NewURLGenerator(testLog, num)

	for ; num < 8; num += 1 {
		url, err := generator.Get()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%d", num)
		assert.Equal(t, url, expected)

		expectedLog := fmt.Sprintf("Generated new short URL: %d\n", num)
		assert.Equal(t, testLog.PrevLog, expectedLog)
	}
}
