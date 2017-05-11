package videoer

import (
	"testing"
	"tribe/parser"
)

func TestFFMPEGImgsToVideo(t *testing.T) {
	parser.ParseFile("../test.json")
	f := CreateFile(parser.Tracker)
	ffmpeg := &FFMPEG{
		Command:   "-f concat -safe 0 -i",
		Input:     f.Path,
		PixFmt:    "yuv420p",
		Vsync:     "vfr",
		Overwrite: true,
	}
	ffmpeg.ImgsToVideo("/Users/jonathansteenbergen/go/src/tribe/output.mp4")
}
