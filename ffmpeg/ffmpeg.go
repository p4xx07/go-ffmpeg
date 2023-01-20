package ffmpeg

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/arguments"
	"github.com/DanielEnricoBotta94/go-ffmpeg/configuration"
	"os/exec"
)

type Ffmpeg struct {
	Arguments     Arguments
	Configuration configuration.Configuration
	Headers       []string
}

type Arguments struct {
	Inputs        arguments.Inputs
	Outputs       arguments.Outputs
	FilterComplex arguments.FilterComplex
	Options       arguments.Options
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

func (f *Ffmpeg) Output(path string) *Ffmpeg {
	f.Arguments.Outputs.Append(path)
	return f
}

func (f *Ffmpeg) Outputs(paths ...string) *Ffmpeg {
	for _, p := range paths {
		f.Arguments.Outputs.Append(p)
	}
	return f
}

func (f *Ffmpeg) Build() []string {
	var args []string
	args = append(args, f.Arguments.Inputs.Build()...)
	args = append(args, f.Arguments.FilterComplex.Build()...)
	args = append(args, f.Arguments.Options.Build()...)
	args = append(args, f.Arguments.Outputs.Build()...)

	return args
}

func (f *Ffmpeg) Run(args []string) error {
	command := exec.Command(f.Configuration.FfmpegPath)

	command.Args = append(command.Args, f.Headers...)
	command.Args = append(command.Args, args...)

	result, err := command.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(result))
		return err
	}

	return nil
}
