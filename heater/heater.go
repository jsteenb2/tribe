package heater

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

func CreateHeatMap() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	cw := imagick.NewPixelWand()

	diameter := uint(75)
	radius := float64(diameter / 2)

	cw.SetColor("transparent")
	mw.NewImage(1000, 1000, cw)

	dw.SetStrokeOpacity(0.5)

	dw.PushDrawingWand()

	cw.SetColor("red")
	dw.SetStrokeColor(cw)
	dw.SetStrokeWidth(4)
	dw.SetStrokeAntialias(true)

	cw.SetColor("red")
	dw.SetFillOpacity(0.5)
	dw.SetFillColor(cw)
	dw.Circle(radius, radius, radius, radius*2)

	mw.DrawImage(dw)
	mw.WriteImage("chart_test.png")
}
