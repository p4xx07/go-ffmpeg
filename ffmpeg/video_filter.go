package ffmpeg

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/arguments"
)

func (f *Ffmpeg) Scale(width int, height int) *Ffmpeg {
	filter := arguments.Filter{
		Name: "scale",
		Args: []arguments.FilterArgs{
			{
				Key:   "width",
				Value: fmt.Sprintf("%d", width),
			},
			{
				Key:   "height",
				Value: fmt.Sprintf("%d", height),
			},
		},
	}
	f.Arguments.VideoFilters.Append(filter)
	return f
}

func (f *Ffmpeg) Fps(fps float64) *Ffmpeg {
	filter := arguments.Filter{
		Name: "fps",
		Args: []arguments.FilterArgs{
			{
				Key:   "fps",
				Value: fmt.Sprintf("%f", fps),
			},
		},
	}
	f.Arguments.VideoFilters.Append(filter)
	return f
}
