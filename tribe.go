package tribe

import (
	"tribe/heater"
	"tribe/parser"
)

func main() {

}

// Runner is the main runner that brings in the files
//  will be run with data receivied from user input on front end
func Runner(video VideoInterface, jsonFilename string) {
	parser.ParseFile(jsonFilename)
	heater.CreateHeatMaps(parser.Tracker, video)
}

// VideoInterface represents interface for giving video information to the inputs
type VideoInterface interface {
	GetHeight() uint
	GetWidth() uint
}
