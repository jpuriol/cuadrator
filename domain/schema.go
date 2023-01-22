package domain

type Schema struct {
	Name        string
	Shifts      map[int]string
	Occupations map[int]string
}

func (s Schema) ShiftName(shiftID int) string {
	return s.Shifts[shiftID]
}

func (s Schema) OcupationName(occupationID int) string {
	return s.Occupations[occupationID]
}
