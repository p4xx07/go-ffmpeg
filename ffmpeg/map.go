package ffmpeg

type Map Ffmpeg

func (fg *Chain) Map(output string) *Map {
	m := Map{
		Arguments:     fg.Arguments,
		Configuration: fg.Configuration,
		Headers:       fg.Headers,
	}
	m.Arguments.FilterGraph.MapChain = append(m.Arguments.FilterGraph.MapChain, output)
	return &m
}

func (m *Map) Map(output string) *Map {
	m.Arguments.FilterGraph.MapChain = append(m.Arguments.FilterGraph.MapChain, output)
	return m
}

func (m *Map) Output(output string) *Builder {
	builder := Builder{
		Arguments:     m.Arguments,
		Configuration: m.Configuration,
		Headers:       m.Headers,
	}
	return builder.Output(output)
}
