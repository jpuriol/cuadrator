package data

import (
    "github.com/jpuriol/cuadrator/domain"
    "os"

	"gopkg.in/yaml.v3"
)

func ReadParticipants() (domain.Participants, error) {
	data, err := os.ReadFile(participantsFile)
	if err != nil {
		return []string{}, err
	}

	var participants domain.Participants
	err = yaml.Unmarshal(data, &participants)
	if err != nil {
		return []string{}, err
	}

	return participants, nil
}
