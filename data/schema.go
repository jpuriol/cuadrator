package data

import (
    "github.com/jpuriol/cuadrator/domain"
    "os"

	"gopkg.in/yaml.v3"
)

func ReadSchema() (domain.Schema, error) {
	rawData, err := os.ReadFile(schemaFile)
	if err != nil {
		return domain.Schema{}, err
	}

	var s domain.Schema
	err = yaml.Unmarshal(rawData, &s)
	if err != nil {
		return domain.Schema{}, err
	}

	return s, nil
}
