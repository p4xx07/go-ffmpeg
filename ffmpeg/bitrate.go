package ffmpeg

import (
	"fmt"
)

func (f *ffmpeg) BitrateVideo(bitrate int) IFfmpeg {
	f.arguments.Options.Append("-b:v", fmt.Sprintf("%d", bitrate))
	return f
}

func (f *ffmpeg) BitrateAudio(bitrate int) IFfmpeg {
	f.arguments.Options.Append("-b:a", fmt.Sprintf("%d", bitrate))
	return f
}
