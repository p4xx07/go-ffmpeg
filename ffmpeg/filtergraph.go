package ffmpeg

import "fmt"

type FilterGraph struct {
	FilterChain []string
	MapChain    []string
}

func (f *Ffmpeg) FilterGraph() *Chain {
	return &Chain{
		Arguments:     f.Arguments,
		Configuration: f.Configuration,
		Headers:       f.Headers,
	}
}

func (fg *FilterGraph) Build() []string {
	result := []string{"-filter_complex"}
	for i, value := range fg.FilterChain {
		if i < len(fg.FilterChain)-1 {
			result = append(result, fmt.Sprintf("%s;", value))
			continue
		}
		result = append(result, fmt.Sprintf("%s", value))
	}

	for _, value := range fg.MapChain {
		result = append(result, "-map")
		result = append(result, fmt.Sprintf("[%s]", value))
	}

	return result
}
