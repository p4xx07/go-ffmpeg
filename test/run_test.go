package test

import (
	"github.com/DanielEnricoBotta94/go-ffmpeg/constants/codec_audio"
	"github.com/DanielEnricoBotta94/go-ffmpeg/constants/codec_video"
	"testing"
)

func Test_Run(t *testing.T) {
	f := GetFfmpeg()
	args := f.
		Input(input).
		Output(output).
		Build()
	err := f.Run(args)

	if err != nil {
		panic(err.Error())
	}
}

func Test_Encode(t *testing.T) {
	f := GetFfmpeg()
	args := f.
		Input(input).
		Scale(500, 500).
		Fps(15).
		Flanger(15).
		CodecVideo(codec_video.LIBX264).
		CodecAudio(codec_audio.AAC).
		Output(output).
		Build()

	err := f.Run(args)

	if err != nil {
		panic(err.Error())
	}
}
