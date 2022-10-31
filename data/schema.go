package data

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Schema struct {
	Name        string
	Shifts      map[int]string
	Occupations map[int]string
}

func readSchema() (Schema, error) {
	rawData, err := os.ReadFile(schemaFile)
	if err != nil {
		return Schema{}, err
	}

	var s Schema
	err = yaml.Unmarshal(rawData, &s)
	if err != nil {
		return Schema{}, err
	}

	return s, nil
}

func (s Schema) shiftName(shiftID int) string {
	return s.Shifts[shiftID]
}

func (s Schema) ocupationName(occupationID int) string {
	return s.Occupations[occupationID]
}
