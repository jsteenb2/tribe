package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)
	ParseFile("../test.json")

	result := Tracker[0]

	assert.Equal("tracker", result.Category, "should be equal")
	assert.Equal(200, result.Statuscode, "status code should be 200")

	var string string
	assert.IsType(string, result.Values.Frame.Timestamp, "expecte tiemstamp to be a string")

	assert.Equal(303, len(Tracker), "Expected 303 tracker data points")

	assert.Equal(11, len(Heartbeat), "Expected 13 tracker data points")
}
