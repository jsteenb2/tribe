package videoer

import (
	"testing"
	"tribe/parser"
)

func TestCreateHeatmapVideo(t *testing.T) {
	video := videoData{Height: 1050, Width: 1680}
	parser.ParseFile("../test.json")
	CreateHeatmapVideo(video, parser.Tracker)
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
