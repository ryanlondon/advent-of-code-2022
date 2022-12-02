package utils

func Filter[T any](arr []T, fn func(T) bool) []T {
	var filteredArr []T

	for _, item := range arr {
		if fn(item) {
			filteredArr = append(filteredArr, item)
		}
	}

	return filteredArr
}
