package helper

func RemoveEmptyEntries(args []string) []string {
	var result []string
	for _, v := range args {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}
