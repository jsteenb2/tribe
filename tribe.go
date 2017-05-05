package tribe

import (
	"tribe/heater"
	"tribe/parser"
)

func main() {

}

func Runner(video VideoInterface, jsonFilename string) {
	parser.ParseFile(jsonFilename)
	heater.CreateHeatMaps(parser.Tracker, video)
}

type VideoInterface interface {
	GetHeight() uint
	GetWidth() uint
}
