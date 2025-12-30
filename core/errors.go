package core

import "fmt"

type ErrParticipantNotFound struct {
	Name string
}

func (e *ErrParticipantNotFound) Error() string {
	return fmt.Sprintf("participant not found: %s", e.Name)
}

type ErrDuplicateAssignment struct {
	ShiftID     int
	Name        string
	Occupations []int
}

func (e *ErrDuplicateAssignment) Error() string {
	return fmt.Sprintf("shift ID %d: duplicate assignment: %s in occupations %v", e.ShiftID, e.Name, e.Occupations)
}
