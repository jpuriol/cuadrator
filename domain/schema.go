package domain

import "fmt"

type Schema struct {
	Name        string
	Shifts      map[int]string
	Occupations map[int]string
}

func (s Schema) ShiftName(shiftID int) string {
	name, ok := s.Shifts[shiftID]
	if !ok {
		return fmt.Sprintf("Unknown Shift %d", shiftID)
	}
	return name
}

func (s Schema) OccupationName(occupationID int) string {
	name, ok := s.Occupations[occupationID]
	if !ok {
		return fmt.Sprintf("Unknown Occupation %d", occupationID)
	}
	return name
}
