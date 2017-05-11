package mediaer

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Probe(path string) Video {
	args := []string{"-v", "error", "-of", "flat=s=_", "-select_streams", "v:0", "-show_entries", "stream=height,width", path}
	cmd := exec.Command(ffprobeLocation(), args...)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	must(cmd.Run())
	resolution := strings.Split(buf.String(), "\n")[:2]

	return Video{
		Height: GetHeight(resolution),
		Width:  GetWidth(resolution),
		Path:   path,
	}
}

func GetHeight(feed []string) uint64 {
	var (
		output uint64
		err    error
	)
	for _, v := range feed {
		if strings.Contains(v, "height=") {
			sizeParam := strings.Split(v, "height=")[1]
			output, err = strconv.ParseUint(sizeParam, 10, 64)
			must(err)
		}
	}
	return output
}

func GetWidth(feed []string) uint64 {
	var (
		output uint64
		err    error
	)
	for _, v := range feed {
		if strings.Contains(v, "width=") {
			sizeParam := strings.Split(v, "width=")[1]
			output, err = strconv.ParseUint(sizeParam, 10, 64)
			must(err)
		}
	}
	return output
}
func ffprobeLocation() string {
	path := "/usr/local/bin/ffprobe"
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

type Video struct {
	Height uint64
	Width  uint64
	Path   string
}
