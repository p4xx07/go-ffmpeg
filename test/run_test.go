package test

import (
	"testing"
)

func Test_Run(t *testing.T) {

	args := ffmpeg.
		Input(input).
		Output(output).
		Build()

	err := ffmpeg.Run(args)
	if err != nil {
		panic(err.Error())
	}
}
