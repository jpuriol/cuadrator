package core

import (
	"fmt"
)

// ValidateSchedule ensures all names in the quadrant are valid participants and no participant is assigned more than once to the same shift.
func ValidateSchedule(q Schedule, p Participants) error {
	for shiftID, shift := range q {
		nameToOccupations := make(map[string][]int)
		for occupationID, groups := range shift {
			for _, group := range groups {
				for _, name := range group {
					nameToOccupations[name] = append(nameToOccupations[name], occupationID)
				}
			}
		}

		for name, occupations := range nameToOccupations {
			if !p.Exists(name) {
				return fmt.Errorf("shift ID %d: %w", shiftID, &ErrParticipantNotFound{Name: name})
			}
			if len(occupations) > 1 {
				return &ErrDuplicateAssignment{ShiftID: shiftID, Name: name, Occupations: occupations}
			}
		}
	}

	return nil
}
