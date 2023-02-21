package ffmpeg

import (
	"fmt"
	"github.com/Paxx-RnD/go-ffmpeg/arguments"
)

func (f *ffmpeg) Scale(width int, height int) IFfmpeg {
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
	f.arguments.VideoFilters.Append(filter)
	return f
}

func (f *ffmpeg) Fps(fps float64) IFfmpeg {
	filter := arguments.Filter{
		Name: "fps",
		Args: []arguments.FilterArgs{
			{
				Key:   "fps",
				Value: fmt.Sprintf("%f", fps),
			},
		},
	}
	f.arguments.VideoFilters.Append(filter)
	return f
}
