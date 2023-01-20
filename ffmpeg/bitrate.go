package ffmpeg

import (
	"fmt"
)

func (f *Ffmpeg) BitrateVideo(bitrate int) *Ffmpeg {
	f.Arguments.Options.Append("-b:v", fmt.Sprintf("%d", bitrate))
	return f
}

func (f *Ffmpeg) BitrateAudio(bitrate int) *Ffmpeg {
	f.Arguments.Options.Append("-b:a", fmt.Sprintf("%d", bitrate))
	return f
}
