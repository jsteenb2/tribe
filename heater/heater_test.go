package heater_test

import (
	"testing"
	"tribe/heater"
)

func TestCreateHeatMap(t *testing.T) {
	coords := eyeData{X: 800.0, Y: 600.0}
	video := videoData{Height: 1000, Width: 1000}
	heater.CreateHeatMap("0", coords, video)
}

type eyeData struct {
	X float64
	Y float64
}

func (e eyeData) GetXCoord() float64 {
	return e.X
}

func (e eyeData) IsFixed() bool {
	return true
}

func (e eyeData) GetYCoord() float64 {
	return e.Y
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
