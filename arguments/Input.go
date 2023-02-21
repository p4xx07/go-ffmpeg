package arguments

type Inputs []input
type input struct {
	value   string
	isLavfi bool
}

func NewInput(value string, isLavfi bool) input {
	return input{
		value:   value,
		isLavfi: isLavfi,
	}
}

func (i *Inputs) Append(input input) *Inputs {
	*i = append(*i, input)
	return i
}

func (i *Inputs) Build() []string {
	var result []string
	for _, v := range *i {
		if v.isLavfi {
			result = append(result, "-f")
			result = append(result, "lavfi")
		}

		result = append(result, "-i")
		result = append(result, v.value)
	}
	return result
}
