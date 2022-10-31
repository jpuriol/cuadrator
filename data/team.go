package data

type Team []string

func (t Team) HasParticipant(name string) bool {

	for _, p := range t {
		if p == name {
			return true
		}
	}

	return false
}
