package ffmpeg

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/arguments"
)

func (f *Ffmpeg) Flanger(delay int) *Ffmpeg {
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
				Value: fmt.Sprintf("%d", delay),
			},
		},
	}
	f.Arguments.AudioFilters.Append(filter)
	return f
}
