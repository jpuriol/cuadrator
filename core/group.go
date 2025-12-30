package core

// Group (Team) is a list of participant names working together.
type Group []string

// HasParticipant checks if a participant is part of the group.
func (g Group) HasParticipant(name string) bool {
	for _, p := range g {
		if p == name {
			return true
		}
	}
	return false
}
