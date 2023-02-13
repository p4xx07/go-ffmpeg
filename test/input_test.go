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
