package core

import (
	"fmt"
	"sort"
)

// Quadrant represents the full schedule, mapping shift IDs to Shift data.
type Quadrant map[int]Shift

// ValidateNames ensures all names in the quadrant are valid participants.
func (q Quadrant) ValidateNames(p Participants) error {

	for shiftID, shift := range q {
		nameFreq := shift.NameFrequency()
		for name := range nameFreq {
			if !p.Exists(name) {
				return fmt.Errorf("name %q on shift ID %d is not in participants", name, shiftID)
			}
		}
	}

	return nil
}

// ValidateShifts ensures no participant is assigned more than once to the same shift.
func (q Quadrant) ValidateShifts() error {
	for shiftID, shift := range q {
		nameFreq := shift.NameFrequency()
		for name, freq := range nameFreq {
			if freq > 1 {
				return fmt.Errorf("participant %q has %d occupations on shift ID %d", name, freq, shiftID)
			}
		}

	}

	return nil
}

// OrderedShiftIDs returns the shift IDs sorted in ascending order.
func (q Quadrant) OrderedShiftIDs() []int {
	var shiftIDs []int
	for k := range q {
		shiftIDs = append(shiftIDs, k)
	}
	sort.Ints(shiftIDs)

	return shiftIDs
}

// Occupation represents a participant's assignment to a specific shift and occupation.
type Occupation struct {
	ShiftID      int // ID of the shift
	OccupationID int // ID of the occupation
}

// GetOccupation returns all assignments for a given participant.
func (q Quadrant) GetOccupation(participantName string) []Occupation {
	var occupations []Occupation

	for _, shiftID := range q.OrderedShiftIDs() {
		for occupationID, occupation := range q[shiftID] {
			for _, team := range occupation {
				if team.HasParticipant(participantName) {
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
