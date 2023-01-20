package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Scale(t *testing.T) {
	expected := fmt.Sprintf("-i %s -vf scale=width=1280:height=720 %s", input, output)

	f := GetFfmpeg()
	args := f.
		Input(input).
		Scale(1280, 720).
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}
}

func Test_Fps(t *testing.T) {
	expected := fmt.Sprintf("-i %s -vf fps=fps=69.000000 %s", input, output)

	f := GetFfmpeg()
	args := f.
		Input(input).
		Fps(69).
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}
}

func Test_ScaleAndFps(t *testing.T) {
	expected := fmt.Sprintf("-i %s -vf scale=width=1280:height=720,fps=fps=69.000000 %s", input, output)

	f := GetFfmpeg()
	args := f.
		Input(input).
		Scale(1280, 720).
		Fps(69).
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}
}
