package domain

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
