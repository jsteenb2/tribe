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

	diameter := uint(500)
	radius := float64(diameter / 2)

	cw.SetColor("transparent")
	mw.NewImage(500, 500, cw)

	dw.PushDrawingWand()

	cw.SetColor("red")
	cw.SetAlpha(0.5)

	dw.SetStrokeColor(cw)
	dw.SetStrokeWidth(1)
	dw.SetStrokeAntialias(true)

	dw.SetFillColor(cw)
	dw.Circle(radius, radius, radius, radius*2)

	mw.DrawImage(dw)
	mw.WriteImage("chart_test.png")
}
