package test

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/configuration"
	"github.com/DanielEnricoBotta94/go-ffmpeg/env"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg"
	"testing"
)

var (
	input  string
	output string
)

func TestMain(m *testing.M) {
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

func PrintPanic(expected any, actual any) {
	panic(fmt.Sprintf("\nexpected: \t%s\nactual: \t%s", expected, actual))
}
