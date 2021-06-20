package maps

import "sort"

func SortedValues(m map[int]string) string {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	concatenated_string := ""
	for _, k := range keys {
		concatenated_string += m[k]
	}
	return concatenated_string
}
