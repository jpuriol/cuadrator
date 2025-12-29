package core

import "sort"

// Shift maps occupation IDs to a list of teams assigned to that occupation.
type Shift map[int][]Team

// NameFrequency counts the number of times each participant appears in the shift.
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

// OrderedOccupationIDs returns the occupation IDs sorted in ascending order.
func (s Shift) OrderedOccupationIDs() []int {
	occupationIDs := make([]int, 0, len(s))

	for k := range s {
		occupationIDs = append(occupationIDs, k)
	}
	sort.Ints(occupationIDs)

	return occupationIDs
}
