package ffprobe

import (
	"strconv"
	"strings"
)

func (p *Probe) GetFps() (float64, error) {
	var videoStream Stream
	for _, stream := range p.Streams {
		if stream.CodecType == "video" {
			videoStream = stream
			break
		}
	}

	rFrameRate := strings.Split(videoStream.RFrameRate, "/")
	numerator, err := strconv.Atoi(rFrameRate[0])
	denominator, err := strconv.Atoi(rFrameRate[1])
	if err != nil {
		return -1, err
	}

	fps := float64(numerator) / float64(denominator)

	return fps, nil
}

func (p *Probe) GetResolution() (int, int) {
	var videoStream Stream
	for _, stream := range p.Streams {
		if stream.CodecType == "video" {
			videoStream = stream
			break
		}
	}

	return videoStream.Width, videoStream.Height
}

func (p *Probe) GetStreamDuration(codecType string) (float64, error) {
	var selected Stream
	for _, stream := range p.Streams {
		if stream.CodecType == codecType {
			selected = stream
			break
		}
	}

	duration, err := strconv.ParseFloat(selected.Duration, 64)
	if err != nil {
		return -1, err
	}

	return duration, err
}
