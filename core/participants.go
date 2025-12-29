package core

type Participants map[string]struct{}

func (p Participants) Exists(name string) bool {
	_, ok := p[name]
	return ok
}
