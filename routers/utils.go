package routers

func del[T interface{}](i int, K []T) []T {
	newArray := make([]T, len(K))
	copy(newArray, K)
	return append(newArray[:i], newArray[i+1:]...)
}
