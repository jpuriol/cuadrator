package core

import "testing"

func TestTeam_HasParticipant(t *testing.T) {
	team := Team{"Alice", "Bob", "Charlie"}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Existing participant", "Alice", true},
		{"Another existing participant", "Bob", true},
		{"Last existing participant", "Charlie", true},
		{"Non-existing participant", "Dave", false},
		{"Empty name", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := team.HasParticipant(tt.input); got != tt.expected {
				t.Errorf("Team.HasParticipant() = %v, want %v", got, tt.expected)
			}
		})
	}
}
