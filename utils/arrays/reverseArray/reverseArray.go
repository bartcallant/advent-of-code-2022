package reverseArray

func Exec[T any](numbers []T) []T {
	newNumbers := make([]T, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}
