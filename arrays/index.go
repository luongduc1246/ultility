package arrays

func ConvertToSliceInterface[T any](t []T) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	return s
}
func ConvertSliceInterfaceToSliceType[T any](elm []interface{}) []T {
	s := make([]T, len(elm))
	for _, v := range elm {
		e, ok := v.(T)
		if ok {
			s = append(s, e)
		}
	}
	return s
}
func ConvertSliceTypeToSliceType[T, E any](elm []E) []T {
	iter := ConvertToSliceInterface[E](elm)
	s := ConvertSliceInterfaceToSliceType[T](iter)
	return s
}

func FindIndex[T comparable](a []T, s T) (index int) {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}
func Contain[T comparable](a []T, s T) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func RemoveIndex[T comparable](a []T, i int) []T {
	c := len(a) - 1
	if i > c {
		return a
	} else {
		a[i] = a[c]
		return a[:c]
	}
}

func RemoveItem[T comparable](a []T, s T) []T {
	i := FindIndex(a, s)
	if i != -1 {
		return RemoveIndex(a, i)
	} else {
		return a
	}
}

func SameItem[T comparable](a []T, b []T) (inter []T) {
	low, high := a, b
	if len(a) > len(b) {
		low = b
		high = a
	}
	for _, l := range low {
		for j, h := range high {
			if l == h {
				inter = append(inter, h)
				RemoveIndex(high, j)
				break
			}
		}
	}
	return
}
