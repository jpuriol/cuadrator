package data

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Participants []string

func readParticipants() (Participants, error) {
	data, err := os.ReadFile(participantsFile)
	if err != nil {
		return []string{}, err
	}

	var participants Participants
	err = yaml.Unmarshal(data, &participants)
	if err != nil {
		return []string{}, err
	}

	return participants, nil
}

func (p Participants) Exists(name string) bool {
	for _, pName := range p {
		if name == pName {
			return true
		}
	}

	return false
}
