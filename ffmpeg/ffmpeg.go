package ffmpeg

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/arguments"
	"github.com/DanielEnricoBotta94/go-ffmpeg/configuration"
	"os/exec"
	"strings"
)

type Ffmpeg struct {
	Arguments     Arguments
	Configuration configuration.Configuration
	Headers       []string
}

type Builder Ffmpeg

type Arguments struct {
	Inputs       arguments.Inputs
	Outputs      arguments.Outputs
	VideoFilters arguments.VideoFilters
	AudioFilters arguments.AudioFilters
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

func (f *Ffmpeg) Output(path string) *Builder {
	f.Arguments.Outputs.Append(path)
	return (*Builder)(f)
}

func (f *Builder) Output(path string) *Builder {
	f.Arguments.Outputs.Append(path)
	return f
}

func (f *Ffmpeg) Outputs(paths ...string) *Builder {
	for _, p := range paths {
		f.Arguments.Outputs.Append(p)
	}
	return (*Builder)(f)
}

func (f *Builder) Build() []string {
	var args []string
	args = append(args, f.Arguments.Inputs.Build()...)
	args = append(args, f.Arguments.VideoFilters.Build()...)
	args = append(args, f.Arguments.AudioFilters.Build()...)
	args = append(args, f.Arguments.Options.Build()...)
	args = append(args, f.Arguments.Outputs.Build()...)

	return args
}

func (f *Ffmpeg) Run(args []string) error {
	command := exec.Command(f.Configuration.FfmpegPath)

	command.Args = append(command.Args, f.Headers...)
	command.Args = append(command.Args, args...)

	command.Args = purifyArgs(command.Args)

	if f.Configuration.IsDebug {
		fmt.Println(strings.Join(command.Args, " "))
	}

	result, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf(fmt.Sprint(err) + ": " + string(result))
		return err
	}

	return nil
}

func purifyArgs(args []string) []string {
	var result []string
	for _, v := range args {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}
