package helper

func TernaryBool(condition bool) int {
	return Ternary[int](condition, 1, 0)
}

func Ternary[T any](condition bool, success T, failure T) T {
	if condition {
		return success
	}
	return failure
}

func TernaryPredicate[T any](condition func() bool, success T, failure T) T {
	if condition() {
		return success
	}
	return failure
}
