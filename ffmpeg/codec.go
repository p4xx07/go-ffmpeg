package ffmpeg

import (
	"github.com/Paxx-RnD/go-ffmpeg/constants/codec_audio"
	"github.com/Paxx-RnD/go-ffmpeg/constants/codec_video"
)

func (f *Ffmpeg) CodecVideo(codec codec_video.CodecVideo) IFfmpeg {
	f.arguments.Options.Append("-c:v", string(codec))
	return f
}

func (f *Ffmpeg) CodecAudio(codec codec_audio.CodecAudio) IFfmpeg {
	f.arguments.Options.Append("-c:a", string(codec))
	return f
}
