package ffmpeg

import (
	"fmt"
	"github.com/DanielEnricoBotta94/go-ffmpeg/constants/pixel_formats"
	"github.com/DanielEnricoBotta94/go-ffmpeg/helper"
)

type Chain Ffmpeg

func (fg *Chain) Format(input string, format pixel_formats.PixelFormat, output string) *Chain {
	chain := fmt.Sprintf("[%s]format=%s[%s]", input, format, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Fps(input string, fps float64, output string) *Chain {
	chain := fmt.Sprintf("[%s]fps=fps=%f[%s]", input, fps, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Scale(input string, width float64, height float64, output string) *Chain {
	chain := fmt.Sprintf("[%s]scale=width=%f:height=%f[%s]", input, width, height, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Reverse(input string, output string) *Chain {
	chain := fmt.Sprintf("[%s]reverse[%s]", input, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Trim(input string, start float64, end float64, output string) *Chain {
	chain := fmt.Sprintf("[%s]trim=start=%f:end=%f,setpts=PTS-STARTPTS[%s]", input, start, end, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Split(input string, outputs ...string) *Chain {
	chain := fmt.Sprintf("[%s]split=n=%d", input, len(outputs))
	for _, output := range outputs {
		chain += fmt.Sprintf("[%s]", output)
	}
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) ATrim(input string, start float64, end float64, output string) *Chain {
	chain := fmt.Sprintf("[%s]atrim=start=%f:end=%f,asetpts=PTS-STARTPTS[%s]", input, start, end, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) AlphaExtract(input string, output string) *Chain {
	chain := fmt.Sprintf("[%s]alphaextract[%s]", input, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) AlphaMerge(input string, mask string, shortest bool, output string) *Chain {
	short := helper.TernaryBool(shortest)
	chain := fmt.Sprintf("[%s][%s]alphamerge=shortest=%d[%s]", input, mask, short, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Concat(inputs []string, videoEnable bool, audioEnable bool, output string) *Chain {
	toConcat := make([]string, len(inputs))
	for i, input := range inputs {
		toConcat[i] = fmt.Sprintf("[%s]", input)
	}

	vFlag := helper.TernaryBool(videoEnable)
	aFlag := helper.TernaryBool(audioEnable)

	chain := fmt.Sprintf("%sconcat=v=%d:a=%d[%s]", toConcat, vFlag, aFlag, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Overlay(under string, over string, shortest bool, output string) *Chain {
	short := helper.TernaryBool(shortest)
	chain := fmt.Sprintf("[%s][%s]overlay=shortest=%d[%s]", under, over, short, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) ACrossFade(input1 string, input2 string, duration float64, output string) *Chain {
	chain := fmt.Sprintf("[%s][%s]acrossfade=d=%f[%s]", input1, input2, duration, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) ADelay(input string, delay float64, output string) *Chain {
	chain := fmt.Sprintf("[%s]adelay=%f[%s]", input, delay, output)
	fg.Arguments.FilterGraph.FilterChain = append(fg.Arguments.FilterGraph.FilterChain, chain)
	return fg
}
