package tribe

import (
	"strconv"
	"time"
	"tribe/heater"
	"tribe/parser"
)

func main() {

}

func Runner(video VideoInterface) {
	parser.ParseFile("test.json")
	for i, tracker := range parser.Tracker {
		if i%18 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
		heater.CreateHeatMap(strconv.FormatInt(int64(i), 10), tracker, video)
	}
}

type VideoInterface interface {
	GetHeight() uint
	GetWidth() uint
}
