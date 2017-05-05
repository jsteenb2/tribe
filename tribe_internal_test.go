package tribe

import (
	"strconv"
	"testing"
	"tribe/heater"
	"tribe/parser"
)

func TestMain(t *testing.T) {
	video := videoData{Height: 1050, Width: 1680}
	parser.ParseFile("test.json")
	for i, tracker := range parser.Tracker {
		heater.CreateHeatMap(strconv.FormatInt(int64(i), 10), tracker, video)
	}
}

type videoData struct {
	Height int
	Width  int
}

func (v videoData) GetHeight() uint {
	return uint(v.Height)
}

func (v videoData) GetWidth() uint {
	return uint(v.Width)
}

// /Users/jonathansteenbergen/async/eye_tracker/imgs/one_min_test.mp4
