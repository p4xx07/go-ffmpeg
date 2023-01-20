package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_OneInputOneOutput(t *testing.T) {
	expected := fmt.Sprintf("-i %s %s", input, output)

	f := GetFfmpeg()
	args := f.
		Input(input).
		Output(output).
		Build()

	if strings.Join(args, " ") != expected {
		panic("failed")
	}
}

func Test_TwoInputTwoOutput(t *testing.T) {
	expected := fmt.Sprintf("-i %s -i %s %s %s", input, input, output, "output2.mp4")

	f := GetFfmpeg()
	args := f.
		Input(input).
		Input(input).
		Output(output).
		Output("output2.mp4").
		Build()

	if strings.Join(args, " ") != expected {
		panic("failed")
	}
}

func Test_TwoInputs(t *testing.T) {
	expected := fmt.Sprintf("-i %s -i %s %s", input, input, output)

	f := GetFfmpeg()
	args := f.
		Inputs(input, input).
		Output(output).
		Build()

	if strings.Join(args, " ") != expected {
		panic("failed")
	}
}

func Test_TwoOutputs(t *testing.T) {
	expected := fmt.Sprintf("-i %s %s %s", input, output, output)

	f := GetFfmpeg()
	args := f.
		Input(input).
		Outputs(output, output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}

}
