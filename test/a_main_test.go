package test

import (
	"github.com/DanielEnricoBotta94/go-ffmpeg/configuration"
	ffmpeg2 "github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg"
	"os"
	"testing"
)

var ffmpeg ffmpeg2.Ffmpeg

const input = "/Users/danielenricobotta/Work/ffmpeg/intro.mp4"
const output = "/Users/danielenricobotta/Work/ffmpeg/output.mp4"

func TestMain(m *testing.M) {
	ffmpeg = ffmpeg2.Ffmpeg{
		Arguments:     ffmpeg2.Arguments{},
		Configuration: configuration.Configuration{},
		Headers: []string{
			"-y",
			"-hide_banner",
			"-loglevel", "info",
		},
	}

	ffmpeg.Configuration.Init()
	code := m.Run()
	os.Exit(code)
}
