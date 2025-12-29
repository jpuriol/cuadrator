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

type FullData struct {
	Quadrant     core.Quadrant
	Participants core.Participants
	Schema       core.Schema
}

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

func GeneratePDF() error {
	d, err := LoadAll()
	if err != nil {
		return err
	}

	return adapters.WritePDF(d.Quadrant, d.Participants, d.Schema)
}
