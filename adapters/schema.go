package adapters

import (
	"github.com/jpuriol/cuadrator/core"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadSchema(filename string) (core.Schema, error) {
	rawData, err := os.ReadFile(filename)
	if err != nil {
		return core.Schema{}, err
	}

	var s core.Schema
	err = yaml.Unmarshal(rawData, &s)
	if err != nil {
		return core.Schema{}, err
	}

	return s, nil
}
