package arguments

type Inputs []string

func (i *Inputs) Append(path string) *Inputs {
	*i = append(*i, path)
	return i
}

func (i *Inputs) Build() []string {
	var result []string
	for _, v := range *i {
		result = append(result, "-i")
		result = append(result, v)
	}
	return result
}
