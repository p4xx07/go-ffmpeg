package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Flanger(t *testing.T) {
	f := GetFfmpeg()

	expected := fmt.Sprintf("-i %s -af flanger=delay=15.000000 %s", input, output)

	args := f.
		Input(input).
		Flanger(15).
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}
}
