package core

import "testing"

func TestParticipants_Exists(t *testing.T) {
	p := Participants{
		"Alice": {},
		"Bob":   {},
	}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Existing participant", "Alice", true},
		{"Another existing participant", "Bob", true},
		{"Non-existing participant", "Charlie", false},
		{"Empty name", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Exists(tt.input); got != tt.expected {
				t.Errorf("Participants.Exists() = %v, want %v", got, tt.expected)
			}
		})
	}
}
