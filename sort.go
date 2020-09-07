package applytmpl

import (
	"sort"
)

type SortableStrings []string

func (sa SortableStrings) Len() int {
	return len(sa)
}

func (sa SortableStrings) Less(i, j int) bool {
	return sa[i] < sa[j]
}

func (sa SortableStrings) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

func NewSortedKeys(data map[string]string) []string {
	keys := SortableStrings{}
	for key := range data {
		keys = append(keys, key)
	}
	return keys.Sort()
}

func (sa SortableStrings) Sort() []string {
	sort.Sort(sa)
	return []string(sa)
}
