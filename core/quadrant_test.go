package core

import (
	"reflect"
	"testing"
)

func TestQuadrant_ValidateNames(t *testing.T) {
	p := Participants{"Alice": {}, "Bob": {}}

	tests := []struct {
		name    string
		quad    Quadrant
		wantErr bool
	}{
		{
			"All names exist",
			Quadrant{1: Shift{10: []Team{{"Alice", "Bob"}}}},
			false,
		},
		{
			"Name does not exist",
			Quadrant{1: Shift{10: []Team{{"Alice", "Charlie"}}}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.quad.ValidateNames(p); (err != nil) != tt.wantErr {
				t.Errorf("Quadrant.ValidateNames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuadrant_ValidateShifts(t *testing.T) {
	tests := []struct {
		name    string
		quad    Quadrant
		wantErr bool
	}{
		{
			"No duplicates in same shift",
			Quadrant{1: Shift{10: []Team{{"Alice"}}, 20: []Team{{"Bob"}}}},
			false,
		},
		{
			"Duplicate name in same shift (different occupations)",
			Quadrant{1: Shift{10: []Team{{"Alice"}}, 20: []Team{{"Alice"}}}},
			true,
		},
		{
			"Duplicate name in same shift (same occupation, different team)",
			Quadrant{1: Shift{10: []Team{{"Alice"}, {"Alice"}}}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.quad.ValidateShifts(); (err != nil) != tt.wantErr {
				t.Errorf("Quadrant.ValidateShifts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuadrant_OrderedShiftIDs(t *testing.T) {
	q := Quadrant{
		3: Shift{},
		1: Shift{},
		2: Shift{},
	}

	expected := []int{1, 2, 3}
	got := q.OrderedShiftIDs()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Quadrant.OrderedShiftIDs() = %v, want %v", got, expected)
	}
}

func TestQuadrant_GetOccupation(t *testing.T) {
	q := Quadrant{
		1: Shift{10: []Team{{"Alice", "Bob"}}, 20: []Team{{"Charlie"}}},
		2: Shift{10: []Team{{"Alice"}}},
	}

	expected := []Occupation{
		{ShiftID: 1, OccupationID: 10},
		{ShiftID: 2, OccupationID: 10},
	}

	got := q.GetOccupation("Alice")

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Quadrant.GetOccupation() = %v, want %v", got, expected)
	}

	gotEmpty := q.GetOccupation("Dave")
	if len(gotEmpty) != 0 {
		t.Errorf("Quadrant.GetOccupation() for non-existing participant = %v, want empty", gotEmpty)
	}
}
