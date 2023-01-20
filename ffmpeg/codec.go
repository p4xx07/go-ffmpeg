package ffmpeg

import (
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg/codec_audio"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg/codec_video"
)

func (f *Ffmpeg) CodecVideo(codec codec_video.CodecVideo) *Ffmpeg {
	f.Arguments.Options.Append("-c:v", string(codec))
	return f
}

func (f *Ffmpeg) CodecAudio(codec codec_audio.CodecAudio) *Ffmpeg {
	f.Arguments.Options.Append("-c:a", string(codec))
	return f
}
