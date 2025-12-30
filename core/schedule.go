package core

import (
	"sort"
)

// Schedule (Quadrant) represents the full schedule, mapping shift IDs to Shift data.
type Schedule map[int]Shift

// Shift maps occupation IDs to a list of groups assigned to that occupation.
type Shift map[int][]Group

// Occupation represents a participant's assignment to a specific shift and occupation.
type Occupation struct {
	ShiftID      int // ID of the shift
	OccupationID int // ID of the occupation
}

// NameFrequency counts the number of times each participant appears in the shift.
func (s Shift) NameFrequency() map[string]int {
	frequency := make(map[string]int)

	for _, occupation := range s {
		for _, group := range occupation {
			for _, name := range group {
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

// OrderedShiftIDs returns the shift IDs sorted in ascending order.
func (s Schedule) OrderedShiftIDs() []int {
	var shiftIDs []int
	for k := range s {
		shiftIDs = append(shiftIDs, k)
	}
	sort.Ints(shiftIDs)

	return shiftIDs
}

// GetAssignments returns all assignments for a given participant.
func (s Schedule) GetAssignments(participantName string) []Occupation {
	var occupations []Occupation

	for _, shiftID := range s.OrderedShiftIDs() {
		for occupationID, occupation := range s[shiftID] {
			for _, group := range occupation {
				if group.HasParticipant(participantName) {
					occupations = append(occupations, Occupation{
						ShiftID:      shiftID,
						OccupationID: occupationID,
					})
				}
			}
		}
	}

	return occupations
}
