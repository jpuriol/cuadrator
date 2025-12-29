package domain

import (
	"fmt"
	"sort"
)

type Quadrant map[int]Shift

// ValidateNames checks that all the names in quadrant belong to participants
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

// ValidateShifts checks that the same person is not twice on the same shift
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

func (q Quadrant) OrderedShiftIDs() []int {
	var shiftIDs []int
	for k := range q {
		shiftIDs = append(shiftIDs, k)
	}
	sort.Ints(shiftIDs)

	return shiftIDs
}

type Occupation struct {
	ShiftID      int
	OccupationID int
}

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
