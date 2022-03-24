package info

import (
	"os"

	"gopkg.in/yaml.v3"
)

type schema struct {
	Shifts      map[int]string
	Occupations map[int]string
}

func getShiftName(shiftID int) string {
	rawData, _ := os.ReadFile(schemaFile)

	var s schema
	yaml.Unmarshal(rawData, &s)

	return s.Shifts[shiftID]
}

func getOcupationName(occupationID int) string {
	rawData, _ := os.ReadFile(schemaFile)

	var s schema
	yaml.Unmarshal(rawData, &s)

	return s.Occupations[occupationID]
}
