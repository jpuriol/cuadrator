package core

import (
	"reflect"
	"testing"
)

func TestShift_NameFrequency(t *testing.T) {
	shift := Shift{
		1: []Team{{"Alice", "Bob"}, {"Charlie"}},
		2: []Team{{"Alice", "Dave"}},
	}

	expected := map[string]int{
		"Alice":   2,
		"Bob":     1,
		"Charlie": 1,
		"Dave":    1,
	}

	got := shift.NameFrequency()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Shift.NameFrequency() = %v, want %v", got, expected)
	}
}

func TestShift_OrderedOccupationIDs(t *testing.T) {
	shift := Shift{
		10: []Team{{"Alice"}},
		2:  []Team{{"Bob"}},
		5:  []Team{{"Charlie"}},
	}

	expected := []int{2, 5, 10}
	got := shift.OrderedOccupationIDs()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Shift.OrderedOccupationIDs() = %v, want %v", got, expected)
	}
}
