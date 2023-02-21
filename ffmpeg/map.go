package ffmpeg

import "github.com/Paxx-RnD/go-ffmpeg/arguments"

type IMap interface {
	Map(output string) IMap
	Output(output string) IBuilder
}

type Map ffmpeg

func (fg *Chain) Map(output string) IMap {
	m := Map{
		arguments:     fg.arguments,
		configuration: fg.configuration,
		headers:       fg.headers,
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
		configuration: m.configuration,
		headers:       m.headers,
	}
	builder.arguments.Output = arguments.Outputs(output)
	return &builder
}
