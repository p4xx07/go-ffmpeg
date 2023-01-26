package ffmpeg

import (
	"fmt"
	"github.com/Paxx-RnD/go-ffmpeg/helper"
	"os/exec"
	"strings"
)

type Builder Ffmpeg

func (f *Ffmpeg) Output(path string) *Builder {
	f.arguments.Outputs.Append(path)
	return (*Builder)(f)
}

func (f *Builder) Output(path string) *Builder {
	f.arguments.Outputs.Append(path)
	return f
}

func (f *Ffmpeg) Outputs(paths ...string) *Builder {
	for _, p := range paths {
		f.arguments.Outputs.Append(p)
	}
	return (*Builder)(f)
}

func (f *Builder) Build() []string {
	var args []string
	args = append(args, f.arguments.Inputs.Build()...)
	if len(f.arguments.FilterGraph.FilterChain) == 0 {
		args = append(args, f.arguments.VideoFilters.Build()...)
		args = append(args, f.arguments.AudioFilters.Build()...)
	} else {
		args = append(args, f.arguments.FilterGraph.Build()...)
	}
	args = append(args, f.arguments.Options.Build()...)
	args = append(args, f.arguments.Outputs.Build()...)

	return args
}

func (f *Ffmpeg) Run(args []string) error {
	command := exec.Command(f.Configuration.FfmpegPath)

	command.Args = append(command.Args, f.Headers...)
	command.Args = append(command.Args, args...)

	command.Args = helper.RemoveEmptyEntries(command.Args)

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
