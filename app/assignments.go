package app

import (
	"context"
	"github.com/jpuriol/cuadrator/core"
)

type ParticipantAssignments struct {
	Name        string
	Assignments []core.Occupation
	Schema      core.Schema
}

func (s *Service) GetAssignments(ctx context.Context, participantName string) (ParticipantAssignments, error) {
	schedule, err := s.store.LoadSchedule(ctx)
	if err != nil {
		return ParticipantAssignments{}, err
	}

	schema, err := s.store.LoadSchema(ctx)
	if err != nil {
		return ParticipantAssignments{}, err
	}

	assignments := schedule.GetAssignments(participantName)

	return ParticipantAssignments{
		Name:        participantName,
		Assignments: assignments,
		Schema:      schema,
	}, nil
}

func (s *Service) GetAllParticipants(ctx context.Context) (core.Participants, error) {
	return s.store.LoadParticipants(ctx)
}

func (s *Service) GetSchema(ctx context.Context) (core.Schema, error) {
	return s.store.LoadSchema(ctx)
}

func (s *Service) GetSchedule(ctx context.Context) (core.Schedule, error) {
	return s.store.LoadSchedule(ctx)
}
