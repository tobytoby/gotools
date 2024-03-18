package utils

import "sort"

func ResetWeight(items map[any]int) map[any]int {
	if len(items) == 0 {
		return nil
	}
	t := make([]struct {
		item   any
		weight int
	}, 0)
	for item, weight := range items {
		t = append(t, struct {
			item   any
			weight int
		}{item: item, weight: weight})
	}
	sort.SliceStable(t, func(i, j int) bool {
		return t[i].weight < t[j].weight
	})
	for i, item := range t {
		items[item] = i
	}
	return items
}
