package test

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/configuration"
	"github.com/DanielEnricoBotta94/go-ffmpeg/env"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffprobe"
	"os"
	"testing"
)

var (
	input  string
	output string
)

func TestMain(m *testing.M) {
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "test")
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

func GetFfprobe() *ffprobe.Ffprobe {
	f := ffprobe.Ffprobe{
		Configuration: configuration.GetConfiguration(),
	}

	input = f.Configuration.TestInput
	output = f.Configuration.TestOutput

	return &f
}

func PrintPanic(expected any, actual any) {
	panic(fmt.Sprintf("\nexpected: \t%s\nactual: \t%s", expected, actual))
}
