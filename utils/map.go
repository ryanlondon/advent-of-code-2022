package utils

func Map[T, M any](arr []T, fn func(T) M) []M {
	var mappedArr []M

	for _, item := range arr {
		mappedItem := fn(item)
		mappedArr = append(mappedArr, mappedItem)
	}

	return mappedArr
}
