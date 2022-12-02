package utils

func Reduce[T, M any](arr []T, fn func(M, T) M, acc M) M {
	for _, item := range arr {
		acc = fn(acc, item)
	}

	return acc
}
