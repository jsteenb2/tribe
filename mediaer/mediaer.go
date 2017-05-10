package mediaer

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func Probe(filename string) Video {
	args := []string{"-v", "error", "-of", "flat=s=_", "-select_streams", "v:0", "-show_entries", "stream=height,width", filename}
	cmd := exec.Command("/usr/local/bin/ffprobe", args...)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	must(cmd.Run())
	resolution := strings.Split(buf.String(), "\n")[:2]

	return Video{
		Height: GetSize("height", resolution),
		Width:  GetSize("width", resolution),
	}
}

func GetSize(parameter string, feed []string) uint64 {
	var (
		output uint64
		err    error
	)
	for _, v := range feed {
		if strings.Contains(v, parameter+"=") {
			sizeParam := strings.Split(v, parameter+"=")[1]
			output, err = strconv.ParseUint(sizeParam, 10, 64)
			must(err)
		}
	}
	return output
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Video struct {
	Height uint64
	Width  uint64
}

func (v Video) GetHeight() uint {
	return uint(v.Height)
}

func (v Video) GetWidth() uint {
	return uint(v.Width)
}
