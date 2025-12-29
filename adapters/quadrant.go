package adapters

import (
	"github.com/jpuriol/cuadrator/core"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadQuadrant reads the quadrant data from a YAML file.
func ReadQuadrant(filename string) (core.Quadrant, error) {
	rawData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data core.Quadrant
	err = yaml.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
