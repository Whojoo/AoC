package shared

func Project[K any, V any](arr []K, f func(K, int) V) []V {
	result := make([]V, len(arr))
	for i, v := range arr {
		result[i] = f(v, i)
	}
	return result
}

func Filter[K any](objs []K, f func(K) bool) []K {
	var result []K
	for _, obj := range objs {
		if f(obj) {
			result = append(result, obj)
		}
	}

	return result
}

func First[K any](objs []K, f func(K) bool) (result *K, found bool) {
	for _, obj := range objs {
		if f(obj) {
			return &obj, true
		}
	}
	return nil, false
}
