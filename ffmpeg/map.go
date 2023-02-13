package ffmpeg

import "github.com/Paxx-RnD/go-ffmpeg/arguments"

type Map Ffmpeg

func (fg *Chain) Map(output string) *Map {
	m := Map{
		arguments:     fg.arguments,
		Configuration: fg.Configuration,
		Headers:       fg.Headers,
	}
	m.arguments.FilterGraph.MapChain = append(m.arguments.FilterGraph.MapChain, output)
	return &m
}

func (m *Map) Map(output string) *Map {
	m.arguments.FilterGraph.MapChain = append(m.arguments.FilterGraph.MapChain, output)
	return m
}

func (m *Map) Output(output string) *Builder {
	builder := Builder{
		arguments:     m.arguments,
		Configuration: m.Configuration,
		Headers:       m.Headers,
	}
	builder.arguments.Output = arguments.Outputs(output)
	return &builder
}
