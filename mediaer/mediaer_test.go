package mediaer_test

import (
	"testing"
	"tribe/mediaer"
)

func TestProbe(t *testing.T) {
	video := mediaer.Probe("/Users/jonathansteenbergen/go/src/tribe/output.mp4")
	if video.Height != 1050 {
		t.Errorf("expected height to be `1050` got %d", video.Height)
	}

	if video.Width != 1680 {
		t.Errorf("expected width to be `1680` got %d", video.Width)
	}
}

func TestGetSize(t *testing.T) {
	sample := []string{"streams_stream_0_width=1680", "streams_stream_0_height=1050"}

	height := mediaer.GetSize("height", sample)
	if height != 1050 {
		t.Errorf("expected height to be `1050` got %d", height)
	}

	width := mediaer.GetSize("width", sample)
	if width != 1680 {
		t.Errorf("expected width to be `1680` got %d", width)
	}
}
