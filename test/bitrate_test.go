package test

import (
	"fmt"
	"github.com/Paxx-RnD/go-ffmpeg/constants/common_bitrates"
	"strings"
	"testing"
)

func Test_Bitrate(t *testing.T) {
	f := GetFfmpeg()
	expected := fmt.Sprintf("-i %s -b:v %d -b:a %d %s", input, common_bitrates.VideoBitrate100K, common_bitrates.AudioBitrate128K, output)

	args := f.Input(input).
		BitrateVideo(common_bitrates.VideoBitrate100K).
		BitrateAudio(common_bitrates.AudioBitrate128K).
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}
}
