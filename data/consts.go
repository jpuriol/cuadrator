package data

import "github.com/jpuriol/cuadrator/domain"

const (
	participantsFile = "participants.yaml"
	schemaFile       = "schema.yaml"
	quadrantFile     = "quadrant.yaml"
)

type FullData struct {
	Quadrant     domain.Quadrant
	Participants domain.Participants
	Schema       domain.Schema
}

func LoadAll() (*FullData, error) {
	q, err := ReadQuadrant()
	if err != nil {
		return nil, err
	}

	p, err := ReadParticipants()
	if err != nil {
		return nil, err
	}

	s, err := ReadSchema()
	if err != nil {
		return nil, err
	}

	return &FullData{
		Quadrant:     q,
		Participants: p,
		Schema:       s,
	}, nil
}
