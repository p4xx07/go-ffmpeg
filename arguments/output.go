package arguments

type Outputs []string

func (i *Outputs) Append(path string) *Outputs {
	*i = append(*i, path)
	return i
}

func (i *Outputs) Build() []string {
	return *i
}
