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

func (p *Probe) GetResolution() (width int, height int) {
	var videoStream Stream
	for _, stream := range p.Streams {
		if stream.CodecType == "video" {
			videoStream = stream
			break
		}
	}

	return videoStream.Width, videoStream.Height
}

func (p *Probe) GetStreamDuration() (float64, error) {
	duration, err := strconv.ParseFloat(p.Format.Duration, 64)
	if err != nil {
		return -1, err
	}

	return duration, err
}
