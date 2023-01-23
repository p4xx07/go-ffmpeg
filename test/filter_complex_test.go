package test

/*

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/constants/codec_video"
	"testing"
)

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/constants/codec_audio"
	"github.com/DanielEnricoBotta94/go-ffmpeg/constants/codec_video"
	"strings"
	"testing"
)

func Test_FilterComplex(t *testing.T) {
	expected := fmt.Sprintf("-i %s -filter_complex  -c:v libx264 -c:a aac %s", input, output)

	f := GetFfmpeg()
	args := f.
		Input(input).
		CodecVideo(codec_video.LIBX264).
		CodecAudio(codec_audio.AAC).
		FilterGraph().
		Fps("0:v", 15, "fps1").
		Scale("fps1", 100, 100, "scale").
		Map("scale").
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}

}
*/
