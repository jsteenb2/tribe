package heater

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

// CreateHeatMap creates a png image with coordinates from viewers gaze, fixation, and the the Video details
func CreateHeatMap(filename string, coords ViewerInterface, video VideoInterface) {
	imagick.Initialize()
	defer imagick.Terminate()

	dw := buildDrawing(coords)

	mw := buildImage(video)
	mw.DrawImage(dw)
	mw.WriteImage("~/go/src/tribe/heater/imgs/" + filename + ".png")
}

func buildDrawing(coords ViewerInterface) *imagick.DrawingWand {
	radius := 35.0
	dw := imagick.NewDrawingWand()
	dw.PushDrawingWand()
	dw.SetStrokeColor(setColor(coords))
	dw.SetStrokeWidth(1)
	dw.SetStrokeAntialias(true)
	dw.SetStrokeLineCap(imagick.LINE_CAP_ROUND)
	dw.SetStrokeLineJoin(imagick.LINE_JOIN_ROUND)
	dw.SetFillColor(setColor(coords))
	dw.RoundRectangle(coords.GetXCoord(), coords.GetYCoord(), 2*radius+coords.GetXCoord(), 2*radius+coords.GetYCoord(), 2*radius, 2*radius)

	return dw
}

func buildImage(video VideoInterface) *imagick.MagickWand {
	mw := imagick.NewMagickWand()
	cw := imagick.NewPixelWand()
	cw.SetColor("transparent")
	mw.NewImage(video.GetWidth(), video.GetHeight(), cw)
	return mw
}

func setColor(coords ViewerInterface) *imagick.PixelWand {
	cw := imagick.NewPixelWand()
	if coords.IsFixed() {
		cw.SetColor("rgb(255,0,0)")
	} else {
		cw.SetColor("rgb(147, 21, 136)")
	}
	cw.SetAlpha(0.5)
	return cw
}

// ViewerInterface is struct with info about an image file
type ViewerInterface interface {
	GetXCoord() float64
	GetYCoord() float64
	IsFixed() bool
}

// VideoInterface is the interface that maps to an input video file
type VideoInterface interface {
	GetHeight() uint
	GetWidth() uint
}
