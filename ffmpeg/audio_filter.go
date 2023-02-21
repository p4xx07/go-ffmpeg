package ffmpeg

import (
	"fmt"
	"github.com/Paxx-RnD/go-ffmpeg/arguments"
)

func (f *ffmpeg) Flanger(delay float64) IFfmpeg {
	if delay < 0 {
		delay = 0
	} else if delay > 30 {
		delay = 30
	}

	filter := arguments.Filter{
		Name: "flanger",
		Args: []arguments.FilterArgs{
			{
				Key:   "delay",
				Value: fmt.Sprintf("%f", delay),
			},
		},
	}
	f.arguments.AudioFilters.Append(filter)
	return f
}
