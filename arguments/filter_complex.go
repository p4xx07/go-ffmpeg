package arguments

type FilterComplex []string

func (i *FilterComplex) Append(path string) *FilterComplex {
	*i = append(*i, path)
	return i
}

func (i *FilterComplex) Build() []string {
	return *i
}
