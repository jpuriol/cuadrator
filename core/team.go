package core

// Team is a list of participant names working together.
type Team []string

// HasParticipant checks if a participant is part of the team.
func (t Team) HasParticipant(name string) bool {

	for _, p := range t {
		if p == name {
			return true
		}
	}

	return false
}
