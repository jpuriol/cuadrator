package info

import (
	"os"

	"gopkg.in/yaml.v3"
)

const fichParticipantes = "participants.yaml"

func VerParticipantes() ([]string, error) {
	data, err := os.ReadFile(fichParticipantes)
	if err != nil {
		return []string{}, err
	}

	var participantes []string
	err = yaml.Unmarshal(data, &participantes)
	if err != nil {
		return []string{}, err
	}

	return participantes, nil
}

func AñadirParticipante(p string) error {
	participantes, _ := VerParticipantes()

	participantes = append(participantes, p)

	data, err := yaml.Marshal(participantes)
	if err != nil {
		return err
	}

	err = os.WriteFile(fichParticipantes, data, 0664)
	if err != nil {
		return err
	}

	return nil
}
