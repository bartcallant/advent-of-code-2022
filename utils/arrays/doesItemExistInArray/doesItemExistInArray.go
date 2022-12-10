package doesItemExistInArray

func Exec[T any](
	items []T,
	comparer func(a T) bool,
) bool {
	var result bool = false

	for _, item := range items {
		if comparer(item) {
			result = true

			break
		}
	}

	return result
}
