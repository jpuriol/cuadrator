package adapters

import (
	"github.com/jpuriol/cuadrator/core"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadParticipants(filename string) (core.Participants, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var names []string
	err = yaml.Unmarshal(data, &names)
	if err != nil {
		return nil, err
	}

	participants := make(core.Participants)
	for _, name := range names {
		participants[name] = struct{}{}
	}

	return participants, nil
}
