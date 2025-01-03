package shared

func Project[K any, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))
	for i, v := range arr {
		result[i] = f(v)
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

func First[K any](objs *[]K, f func(K) bool) (*K, bool) {
	for i := 0; i < len(*objs); i++ {
		if f((*objs)[i]) {
			return &(*objs)[i], true
		}
	}

	return nil, false
}

func Any[K any](objs []K, f func(K) bool) bool {
	for _, obj := range objs {
		if f(obj) {
			return true
		}
	}
	return false
}
