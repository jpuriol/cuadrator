package app

import (
	"github.com/jpuriol/cuadrator/adapters"
	"github.com/jpuriol/cuadrator/core"
)

const (
	participantsFile = "participants.yaml"
	schemaFile       = "schema.yaml"
	quadrantFile     = "quadrant.yaml"
)

// FullData holds all the information required to generate a quadrant.
type FullData struct {
	Quadrant     core.Quadrant     // The scheduled shifts and occupations
	Participants core.Participants // The list of available participants
	Schema       core.Schema       // The names and structure of the quadrant
}

// LoadAll reads quadrant, participants, and schema from their respective YAML files.
func LoadAll() (*FullData, error) {
	q, err := adapters.ReadQuadrant(quadrantFile)
	if err != nil {
		return nil, err
	}

	p, err := adapters.ReadParticipants(participantsFile)
	if err != nil {
		return nil, err
	}

	s, err := adapters.ReadSchema(schemaFile)
	if err != nil {
		return nil, err
	}

	return &FullData{
		Quadrant:     q,
		Participants: p,
		Schema:       s,
	}, nil
}

// GeneratePDF loads all data and generates the quadrant PDF.
func GeneratePDF() error {
	d, err := LoadAll()
	if err != nil {
		return err
	}

	return adapters.WritePDF(d.Quadrant, d.Participants, d.Schema)
}
