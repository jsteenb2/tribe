package videoer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"tribe/parser"
)

// CreateHeatmapVideo creates a video fo genreated heatmap imgs with corresponding times
func CreateHeatmapVideo(video VideoInterface, trackerData []parser.View) {
	width := video.GetWidth()
	height := video.GetHeight()
	f := createFile(trackerData)
	fmt.Println(width, height, f.Path)
	createVideo()
	// pass file to ffmpeg command
	// delete the temp text file
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createFile(trackerData []parser.View) InputFile {
	inputFile := InputFile{Path: "/Users/jonathansteenbergen/go/src/tribe/input.txt"}
	f, err := os.Create(inputFile.Path)
	check(err)
	defer f.Close()

	trackerLength := len(trackerData)
	for i, data := range trackerData {
		f.WriteString("file '/Users/jonathansteenbergen/go/src/tribe/heater/imgs/" + strconv.FormatInt(int64(i), 10) + ".png'\n")
		f.WriteString("duration " + getDuration(i, trackerLength, data, trackerData) + "\n")
	}
	f.WriteString("file '/Users/jonathansteenbergen/go/src/tribe/heater/imgs/" + strconv.FormatInt(int64(len(trackerData)-1), 10) + ".png'\n")
	f.Sync()
	return inputFile
}

func getDuration(index, length int, current parser.View, trackerData []parser.View) string {
	if index == length-1 {
		return "0.010"
	}
	next := trackerData[index+1]
	return strconv.FormatFloat((next.Duration()-current.Duration())/1000, 'f', -1, 64)
}

func createVideo() {
	filename := "/Users/jonathansteenbergen/go/src/tribe/videoer/gen_video.sh"
	if err := os.Chmod(filename, 0700); err != nil {
		log.Fatal(err)
	}
	_, err := exec.Command("/bin/sh", filename).Output()
	if err != nil {
		log.Fatal(err)
	}
}

type InputFile struct {
	Path string
}

type VideoInterface interface {
	GetHeight() uint
	GetWidth() uint
}
