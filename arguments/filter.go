package arguments

import "fmt"

type Filter struct {
	Name string
	Args []FilterArgs
}

type FilterArgs struct {
	Key   string
	Value string
}

type Filters []Filter

func (i *Filters) Append(filter Filter) *Filters {
	*i = append(*i, filter)
	return i
}

func (i *Filters) Build(filterType string) []string {
	if len(*i) == 0 {
		return make([]string, 0)
	}

	result := []string{filterType}
	var filterAccumulator string
	for i, filter := range *i {
		if i > 0 {
			filterAccumulator += ","
		}

		argAccumulator := fmt.Sprintf("%s=", filter.Name)
		for j, arg := range filter.Args {
			if j > 0 {
				argAccumulator += ":"
			}

			argAccumulator += fmt.Sprintf("%s=%s", arg.Key, arg.Value)
		}
		filterAccumulator += argAccumulator
	}

	result = append(result, filterAccumulator)
	return result
}

type VideoFilters Filters

func (vf *VideoFilters) Build() []string {
	var f = (*Filters)(vf)
	return f.Build("-vf")
}

func (vf *VideoFilters) Append(filter Filter) *Filters {
	var f = (*Filters)(vf)
	return f.Append(filter)
}

type AudioFilters Filters

func (af *AudioFilters) Build() []string {
	var f = (*Filters)(af)
	return f.Build("-af")
}

func (af *AudioFilters) Append(filter Filter) *Filters {
	var f = (*Filters)(af)
	return f.Append(filter)
}
