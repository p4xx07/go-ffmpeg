package test

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/configuration"
	"github.com/DanielEnricoBotta94/go-ffmpeg/env"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg"
	"testing"
)

const input = "/Users/danielenricobotta/Work/ffmpeg/intro.mp4"
const output = "/Users/danielenricobotta/Work/ffmpeg/output.mp4"

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

	return &f
}

func PrintPanic(expected any, actual any) {
	panic(fmt.Sprintf("\nexpected: \t%s\nactual: \t%s", expected, actual))
}
