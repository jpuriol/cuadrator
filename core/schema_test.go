package core

import "testing"

func TestSchema_ShiftName(t *testing.T) {
	s := Schema{
		Shifts: map[int]string{
			1: "Morning",
			2: "Afternoon",
		},
	}

	tests := []struct {
		shiftID  int
		expected string
	}{
		{1, "Morning"},
		{2, "Afternoon"},
		{3, "Unknown Shift 3"},
	}

	for _, tt := range tests {
		if got := s.ShiftName(tt.shiftID); got != tt.expected {
			t.Errorf("Schema.ShiftName(%d) = %q, want %q", tt.shiftID, got, tt.expected)
		}
	}
}

func TestSchema_OccupationName(t *testing.T) {
	s := Schema{
		Occupations: map[int]string{
			10: "Doctor",
			20: "Nurse",
		},
	}

	tests := []struct {
		occID    int
		expected string
	}{
		{10, "Doctor"},
		{20, "Nurse"},
		{30, "Unknown Occupation 30"},
	}

	for _, tt := range tests {
		if got := s.OccupationName(tt.occID); got != tt.expected {
			t.Errorf("Schema.OccupationName(%d) = %q, want %q", tt.occID, got, tt.expected)
		}
	}
}
