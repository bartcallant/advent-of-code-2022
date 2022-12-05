package chunkArrayByChunkSize

func Exec[T any](input []T, chunkSize int) [][]T {
	var chunks [][]T

	for i := 0; i < len(input); i += chunkSize {
		var end = i + chunkSize

		if end > len(input) {
			end = len(input)
		}

		chunks = append(chunks, input[i:end])
	}

	return chunks
}
