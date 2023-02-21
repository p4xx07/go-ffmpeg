package ffmpeg

import "github.com/Paxx-RnD/go-ffmpeg/arguments"

type IMap interface {
	Map(output string) IMap
	Output(output string) IBuilder
}

type Map Ffmpeg

func (fg *Chain) Map(output string) IMap {
	m := Map{
		arguments:     fg.arguments,
		Configuration: fg.Configuration,
		Headers:       fg.Headers,
	}
	m.arguments.FilterGraph.MapChain = append(m.arguments.FilterGraph.MapChain, output)
	return &m
}

func (m *Map) Map(output string) IMap {
	m.arguments.FilterGraph.MapChain = append(m.arguments.FilterGraph.MapChain, output)
	return m
}

func (m *Map) Output(output string) IBuilder {
	builder := Builder{
		arguments:     m.arguments,
		Configuration: m.Configuration,
		Headers:       m.Headers,
	}
	builder.arguments.Output = arguments.Outputs(output)
	return &builder
}
