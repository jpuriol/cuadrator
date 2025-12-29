package core

import "fmt"

// Schema defines the structure of the quadrant, including names for shifts and occupations.
type Schema struct {
	Name         string         // Name of the quadrant
	Title        string         // Title for the PDF
	Subtitle     string         // Subtitle for the PDF
	Shifts       map[int]string // Mapping of shift IDs to names
	Occupations  map[int]string // Mapping of occupation IDs to names
	NoOccupation string         `yaml:"no_occupation"` // Name for when no occupation is assigned
}

// ShiftName returns the name of a shift by its ID.
func (s Schema) ShiftName(shiftID int) string {
	name, ok := s.Shifts[shiftID]
	if !ok {
		return fmt.Sprintf("Unknown Shift %d", shiftID)
	}
	return name
}

// OccupationName returns the name of an occupation by its ID.
func (s Schema) OccupationName(occupationID int) string {
	name, ok := s.Occupations[occupationID]
	if !ok {
		return fmt.Sprintf("Unknown Occupation %d", occupationID)
	}
	return name
}
