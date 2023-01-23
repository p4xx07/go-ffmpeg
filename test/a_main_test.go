package test

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/configuration"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg"
)

const input = "assets/input.mp4"
const output = "assets/output.mp4"

func GetFfmpeg() *ffmpeg.Ffmpeg {
	f := ffmpeg.Ffmpeg{
		Arguments:     ffmpeg.Arguments{},
		Configuration: configuration.Configuration{},
		Headers: []string{
			"-y",
			"-hide_banner",
			"-loglevel", "info",
		},
	}

	f.Configuration.Init()
	return &f
}

func PrintPanic(expected any, actual any) {
	panic(fmt.Sprintf("\nexpected: \t%s\nactual: \t%s", expected, actual))
}
