package videoer

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"tribe/parser"
)

// CreateFile creates input file to use in FFMPEG concat function to build the output video
func CreateFile(trackerData []parser.View) InputFile {
	inputFile := InputFile{Path: "/Users/jonathansteenbergen/go/src/tribe/input.txt"}
	f, err := os.Create(inputFile.Path)
	must(err)
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

// FFMPEG struct for creaeting FFMPEG img to video CLI interface
type FFMPEG struct {
	Command   string  `json:"command"`
	Scale     float64 `json:"scale"`
	PixFmt    string  `json:"pix_fmt"`
	Input     string  `json:"input"`
	Vsync     string  `json:"vsync"`
	Overwrite bool    `json:"overwrite"`
}

// ImgsToVideo creates video from heatmap images
func (ff *FFMPEG) ImgsToVideo(output string) {
	cmd := exec.Command(ffmpegLocation(), ff.buildArguments(output)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	must(cmd.Run())
}

func (ff *FFMPEG) buildArguments(output string) []string {
	args := strings.Split(ff.Command, " ")
	args = append(args, ff.Input)
	if ff.Vsync == "" {
		args = append(args, "-vsync", ff.Vsync)
	}
	if ff.PixFmt == "" {
		args = append(args, "-pix_fmt", ff.PixFmt)
	}
	args = append(args, output, ff.formatOverwrite())
	return args
}

func (ff *FFMPEG) formatOverwrite() string {
	var agreement string
	if ff.Overwrite == true {
		agreement = "-y"
	}
	return agreement
}

func ffmpegLocation() string {
	path := "/usr/local/bin/ffmpeg"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		must(err)
	}

	return path
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// InputFile is representative of file that is created from tracker data
// Relates the images to duration need to map the video to the other video
type InputFile struct {
	Path string
}
