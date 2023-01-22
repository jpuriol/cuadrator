package domain

type Participants []string

func (p Participants) Exists(name string) bool {
	for _, pName := range p {
		if name == pName {
			return true
		}
	}

	return false
}
