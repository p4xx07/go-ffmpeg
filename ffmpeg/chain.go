package ffmpeg

import (
	"fmt"
	"github.com/Paxx-RnD/go-ffmpeg/constants/pixel_formats"
	"github.com/Paxx-RnD/go-helper/helpers/boolean_helper"
)

type IChain interface {
	Format(input string, format pixel_formats.PixelFormat, output string) IChain
	Pad(input string, width int, height int, x int, y int, output string) IChain
	Volume(input string, volume float64, output string) IChain
	Fps(input string, fps float64, output string) IChain
	Scale(input string, width float64, height float64, output string) IChain
	Reverse(input string, output string) IChain
	Trim(input string, start float64, end float64, output string) IChain
	Split(input string, outputs ...string) IChain
	ATrim(input string, start float64, end float64, output string) IChain
	AlphaExtract(input string, output string) IChain
	AlphaMerge(input string, mask string, shortest bool, output string) IChain
	Concat(inputs []string, videoEnable bool, audioEnable bool, output string) IChain
	Overlay(under string, over string, shortest bool, output string) IChain
	ACrossFade(input1 string, input2 string, duration float64, output string) IChain
	ADelay(input string, delay float64, output string) IChain
	Map(output string) IMap
}

type Chain ffmpeg

func (fg *Chain) Format(input string, format pixel_formats.PixelFormat, output string) IChain {
	chain := fmt.Sprintf("[%s]format=%s[%s]", input, format, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Pad(input string, width int, height int, x int, y int, output string) IChain {
	chain := fmt.Sprintf("[%s]pad=width=%d:height=%d:x=%d:y=%d[%s]", input, width, height, x, y, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Volume(input string, volume float64, output string) IChain {
	chain := fmt.Sprintf("[%s]volume=volume=%f[%s]", input, volume, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Fps(input string, fps float64, output string) IChain {
	chain := fmt.Sprintf("[%s]fps=fps=%f[%s]", input, fps, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Scale(input string, width float64, height float64, output string) IChain {
	chain := fmt.Sprintf("[%s]scale=width=%f:height=%f[%s]", input, width, height, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Reverse(input string, output string) IChain {
	chain := fmt.Sprintf("[%s]reverse[%s]", input, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Trim(input string, start float64, end float64, output string) IChain {
	chain := fmt.Sprintf("[%s]trim=start=%f:end=%f,setpts=PTS-STARTPTS[%s]", input, start, end, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Split(input string, outputs ...string) IChain {
	chain := fmt.Sprintf("[%s]split=n=%d", input, len(outputs))
	for _, output := range outputs {
		chain += fmt.Sprintf("[%s]", output)
	}
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) ATrim(input string, start float64, end float64, output string) IChain {
	chain := fmt.Sprintf("[%s]atrim=start=%f:end=%f,asetpts=PTS-STARTPTS[%s]", input, start, end, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) AlphaExtract(input string, output string) IChain {
	chain := fmt.Sprintf("[%s]alphaextract[%s]", input, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) AlphaMerge(input string, mask string, shortest bool, output string) IChain {
	short := boolean_helper.ToInt(shortest)
	chain := fmt.Sprintf("[%s][%s]alphamerge=shortest=%d[%s]", input, mask, short, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Concat(inputs []string, videoEnable bool, audioEnable bool, output string) IChain {
	toConcat := make([]string, len(inputs))
	for i, input := range inputs {
		toConcat[i] = fmt.Sprintf("[%s]", input)
	}

	vFlag := boolean_helper.ToInt(videoEnable)
	aFlag := boolean_helper.ToInt(audioEnable)

	chain := fmt.Sprintf("%sconcat=v=%d:a=%d[%s]", toConcat, vFlag, aFlag, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) Overlay(under string, over string, shortest bool, output string) IChain {
	short := boolean_helper.ToInt(shortest)
	chain := fmt.Sprintf("[%s][%s]overlay=shortest=%d[%s]", under, over, short, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) ACrossFade(input1 string, input2 string, duration float64, output string) IChain {
	chain := fmt.Sprintf("[%s][%s]acrossfade=d=%f[%s]", input1, input2, duration, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}

func (fg *Chain) ADelay(input string, delay float64, output string) IChain {
	chain := fmt.Sprintf("[%s]adelay=%f[%s]", input, delay, output)
	fg.arguments.FilterGraph.FilterChain = append(fg.arguments.FilterGraph.FilterChain, chain)
	return fg
}
