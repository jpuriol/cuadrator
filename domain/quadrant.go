package domain

import (
	"fmt"
	"sort"
)

type Quadrant map[int]Shift

// CheckNames
// Check that all the names in quadrant belong to participants
func (q Quadrant) CheckNames(p Participants) error {

	for shiftID, shift := range q {
		nameFreq := shift.nameFrequency()
		for name := range nameFreq {
			if !p.Exists(name) {
				return fmt.Errorf("name %q on shift ID %d is no in participants", name, shiftID)
			}
		}
	}

	return nil
}

// CheckShifts
// Check that the same person is not twice on the same shift
func (q Quadrant) CheckShifts() error {
	for shiftID, shift := range q {
		nameFreq := shift.nameFrequency()
		for name, freq := range nameFreq {
			if freq > 1 {
				return fmt.Errorf("participant %q has %d ocuppations on shift ID %d", name, freq, shiftID)
			}
		}

	}

	return nil
}

type Occupation struct {
	ShifID       int
	OccupationID int
}

func (q Quadrant) GetOcupation(participantName string) []Occupation {

	var ocupations []Occupation

	var shiftIDs []int
	for k := range q {
		shiftIDs = append(shiftIDs, k)
	}
	sort.Ints(shiftIDs)

	for _, shiftID := range shiftIDs {
		for occupationID, occupation := range q[shiftID] {
			for _, team := range occupation {
				for _, person := range team {
					if person == participantName {
						ocupations = append(ocupations, Occupation{
							ShifID:       shiftID,
							OccupationID: occupationID,
						})
					}
				}
			}
		}
	}

	return ocupations
}
