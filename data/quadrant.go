package data

import (
    "github.com/jpuriol/cuadrator/domain"
    "os"

	"gopkg.in/yaml.v3"
)

func ReadQuadrant() (domain.Quadrant, error) {
	rawData, err := os.ReadFile(quadrantFile)
	if err != nil {
		return nil, err
	}

	var data domain.Quadrant
	err = yaml.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
