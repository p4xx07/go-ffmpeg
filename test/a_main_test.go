package test

import (
	"fmt"
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
	"github.com/Paxx-RnD/go-ffmpeg/env"
	"github.com/Paxx-RnD/go-ffmpeg/ffmpeg"
	"github.com/Paxx-RnD/go-ffmpeg/ffprobe"
	"os"
	"testing"
)

var (
	input  string
	output string
)

func TestMain(m *testing.M) {
	if os.Getenv("APP_ENV") == "" {
		err := os.Setenv("APP_ENV", "test")
		if err != nil {
			panic("could not set test env")
		}
	}
	env.Load()
	m.Run()
}

func GetFfmpeg() *ffmpeg.Ffmpeg {
	f := ffmpeg.Ffmpeg{
		Configuration: configuration.GetConfiguration(),
		Headers: []string{
			"-y",
			"-hide_banner",
			"-loglevel", "info",
		},
	}

	input = f.Configuration.TestInput
	output = f.Configuration.TestOutput

	return &f
}

func GetFfprobe() ffprobe.IFfprobe {
	configuration := configuration.GetConfiguration()
	f := ffprobe.NewFfProbe(configuration, nil)

	input = configuration.TestInput
	output = configuration.TestOutput

	return f
}

func PrintPanic(expected any, actual any) {
	panic(fmt.Sprintf("\nexpected: \t%s\nactual: \t%s", expected, actual))
}
