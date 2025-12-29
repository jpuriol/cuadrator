package core

import "sort"

type Shift map[int][]Team

func (s Shift) NameFrequency() map[string]int {
	frequency := make(map[string]int)

	for _, occupation := range s {
		for _, team := range occupation {
			for _, name := range team {
				frequency[name]++
			}
		}
	}

	return frequency
}

func (s Shift) OrderedOccupationIDs() []int {
	occupationIDs := make([]int, 0, len(s))

	for k := range s {
		occupationIDs = append(occupationIDs, k)
	}
	sort.Ints(occupationIDs)

	return occupationIDs
}
