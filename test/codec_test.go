package test

import (
	"fmt"
	"github.com/Paxx-RnD/go-ffmpeg/constants/codec_audio"
	"github.com/Paxx-RnD/go-ffmpeg/constants/codec_video"
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
