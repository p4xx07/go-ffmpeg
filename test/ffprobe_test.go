package test

import (
	"fmt"
	"testing"
)

func Test_FfprobeStreams(t *testing.T) {
	expected := 2
	ffprobe := GetFfprobe()
	probe, err := ffprobe.GetStreams(input)
	if err != nil {
		panic(fmt.Sprintf("could not probe %s", err.Error()))
	}

	if len(probe.Streams) != expected {
		PrintPanic(expected, len(probe.Streams))
	}
}

func Test_FfprobeVideoStreams(t *testing.T) {
	expected := 1
	expectedVideo := "video"
	ffprobe := GetFfprobe()
	probe, err := ffprobe.GetVideoStreams(input)
	if err != nil {
		panic(fmt.Sprintf("could not probe %s", err.Error()))
	}

	if len(probe.Streams) != expected {
		PrintPanic(expected, len(probe.Streams))
	}

	if probe.Streams[0].CodecType != expectedVideo {
		PrintPanic(expectedVideo, probe.Streams[0].CodecType)
	}
}

func Test_FfprobeAudioStreams(t *testing.T) {
	expected := 1
	expectedAudio := "audio"
	ffprobe := GetFfprobe()
	probe, err := ffprobe.GetAudioStreams(input)
	if err != nil {
		panic(fmt.Sprintf("could not probe %s", err.Error()))
	}

	if len(probe.Streams) != expected {
		PrintPanic(expected, len(probe.Streams))
	}

	if probe.Streams[0].CodecType != expectedAudio {
		PrintPanic(expectedAudio, probe.Streams[0].CodecType)
	}
}
