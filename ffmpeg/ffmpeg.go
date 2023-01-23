package ffmpeg

import (
	"github.com/Paxx-RnD/go-ffmpeg/arguments"
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
)

type Ffmpeg struct {
	Arguments     Arguments
	Configuration configuration.Configuration
	Headers       []string
}

type Arguments struct {
	Inputs       arguments.Inputs
	Outputs      arguments.Outputs
	VideoFilters arguments.VideoFilters
	AudioFilters arguments.AudioFilters
	FilterGraph  FilterGraph
	Options      arguments.Options
}

func (f *Ffmpeg) Input(path string) *Ffmpeg {
	f.Arguments.Inputs.Append(path)
	return f
}

func (f *Ffmpeg) Inputs(paths ...string) *Ffmpeg {
	for _, p := range paths {
		f.Arguments.Inputs.Append(p)
	}
	return f
}
