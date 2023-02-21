package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Lavfi(t *testing.T) {
	f := GetFfmpeg()

	inputLavfi := "color=c=black:duration=3"
	expected := fmt.Sprintf("-f lavfi -i %s -af flanger=delay=15.000000 %s", inputLavfi, output)

	args := f.
		InputLavfi(input).
		Flanger(15).
		Output(output).
		Build()

	actual := strings.Join(args, " ")
	if actual != expected {
		PrintPanic(expected, actual)
	}
}
