package test

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg/codec_audio"
	"github.com/DanielEnricoBotta94/go-ffmpeg/ffmpeg/codec_video"
	"strings"
	"testing"
)

func Test_Codec(t *testing.T) {
	expected := fmt.Sprintf("-i %s -c:v libx264 -c:a aac %s", input, output)

	f := GetFfmpeg()
	args := f.
		Input(input).
		CodecVideo(codec_video.LIBX264).
		CodecAudio(codec_audio.AAC).
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}

}
