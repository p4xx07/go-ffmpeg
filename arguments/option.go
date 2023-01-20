package arguments

import "fmt"

type Option struct {
	key   string
	value string
}

type Options []Option

func (i *Options) Build() []string {
	var result []string
	for _, o := range *i {
		result = append(result, fmt.Sprintf("%s %s", o.key, o.value))
	}
	return result
}

func (i *Options) Append(key string, value string) {
	*i = append(*i, Option{
		key:   key,
		value: value,
	})
}
