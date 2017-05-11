package heater_test

import (
	"fmt"
	"strconv"
	"testing"
	"tribe/heater"
)

func TestCreateHeatMap(t *testing.T) {
	var coords eyeData
	video := videoData{Height: 1680, Width: 1050}

	for i := int64(0); i < 5; i++ {
		x := float64(i) * 200.0
		y := float64(i) * 200.0
		coords = eyeData{X: x, Y: y}
		fmt.Println(int64(i), strconv.FormatInt(i, 10))
		heater.CreateHeatMap(strconv.FormatInt(i, 10), coords, video)
	}
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
