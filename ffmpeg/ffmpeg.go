package ffmpeg

import (
	"github.com/Paxx-RnD/go-ffmpeg/arguments"
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
	"github.com/Paxx-RnD/go-ffmpeg/constants/codec_audio"
	"github.com/Paxx-RnD/go-ffmpeg/constants/codec_video"
)

type IFfmpeg interface {
	BitrateVideo(bitrate int) IFfmpeg
	BitrateAudio(bitrate int) IFfmpeg
	Flanger(delay float64) IFfmpeg
	CodecVideo(codec codec_video.CodecVideo) IFfmpeg
	CodecAudio(codec codec_audio.CodecAudio) IFfmpeg
	Scale(width int, height int) IFfmpeg
	Fps(fps float64) IFfmpeg
	FilterGraph() IChain
	Input(path string) IFfmpeg
	Inputs(paths ...string) IFfmpeg
	Output(path string) IBuilder
	Run(args []string) error
}

type Ffmpeg struct {
	arguments     Arguments
	Configuration *configuration.Configuration
	Headers       []string
}

type Arguments struct {
	Inputs       arguments.Inputs
	Output       arguments.Outputs
	VideoFilters arguments.VideoFilters
	AudioFilters arguments.AudioFilters
	FilterGraph  FilterGraph
	Options      arguments.Options
}

func (f *Ffmpeg) Input(path string) IFfmpeg {
	f.arguments.Inputs.Append(path)
	return f
}

func (f *Ffmpeg) Inputs(paths ...string) IFfmpeg {
	for _, p := range paths {
		f.arguments.Inputs.Append(p)
	}
	return f
}
