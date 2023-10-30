package domain

import "sort"

type Shift map[int][]Team

func (s Shift) nameFrequency() map[string]int {
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

func (s Shift) OrderedOcuppationIDs() []int {
	ocuppationIDs := make([]int, 0, len(s))

	for k := range s {
		ocuppationIDs = append(ocuppationIDs, k)
	}
	sort.Ints(ocuppationIDs)

	return ocuppationIDs
}
