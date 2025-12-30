package core

import (
	"fmt"
)

// ValidateSchedule ensures all names in the quadrant are valid participants and no participant is assigned more than once to the same shift.
func ValidateSchedule(q Schedule, p Participants) error {
	for shiftID, shift := range q {
		nameFreq := shift.NameFrequency()
		for name, freq := range nameFreq {
			if !p.Exists(name) {
				return fmt.Errorf("name %q on shift ID %d is not in participants: %w", name, shiftID, ErrParticipantNotFound)
			}
			if freq > 1 {
				return fmt.Errorf("participant %q has %d occupations on shift ID %d: %w", name, freq, shiftID, ErrDuplicateAssignment)
			}
		}
	}

	return nil
}
