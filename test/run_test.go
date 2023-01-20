package test

import (
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

func Disabled_Test_Encode(t *testing.T) {
	input := "/Users/danielenricobotta/Work/ffmpeg/input.mp4"

	f := GetFfmpeg()
	args := f.
		Input(input).
		Scale(500, 500).
		Fps(15).
		Flanger(15).
		Output(output).
		Build()

	err := f.Run(args)

	if err != nil {
		panic(err.Error())
	}
}
