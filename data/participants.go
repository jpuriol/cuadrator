package data

import (
	"github.com/jpuriol/cuadrator/domain"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadParticipants() (domain.Participants, error) {
	data, err := os.ReadFile(participantsFile)
	if err != nil {
		return nil, err
	}

	var names []string
	err = yaml.Unmarshal(data, &names)
	if err != nil {
		return nil, err
	}

	participants := make(domain.Participants)
	for _, name := range names {
		participants[name] = struct{}{}
	}

	return participants, nil
}
