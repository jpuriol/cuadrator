package core

// Participants represents a set of participant names.
type Participants map[string]struct{}

// Exists checks if a participant with the given name exists.
func (p Participants) Exists(name string) bool {
	_, ok := p[name]
	return ok
}
