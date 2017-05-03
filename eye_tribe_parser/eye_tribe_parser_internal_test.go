package eyeTribeParser

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	ParseFile("../test.json")
	if Tracker[0].Category != "tracker" {
		t.Error("Expected category to equal tracker")
	}
	if Tracker[0].Statuscode != 200 {
		t.Error("Expected status code to equal 200")
	}
}
