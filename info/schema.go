package info

import (
	"os"

	"gopkg.in/yaml.v3"
)

type schema struct {
	Shifts      map[int]string
	Occupations map[int]string
}

func getSchema() (schema, error) {
	rawData, err := os.ReadFile(schemaFile)
	if err != nil {
		return schema{}, err
	}

	var s schema
	err = yaml.Unmarshal(rawData, &s)
	if err != nil {
		return schema{}, err
	}

	return s, nil
}

func (s schema) shiftName(shiftID int) string {
	return s.Shifts[shiftID]
}

func (s schema) ocupationName(occupationID int) string {
	return s.Occupations[occupationID]
}
