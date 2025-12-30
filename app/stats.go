package app

import (
	"context"
	"github.com/jpuriol/cuadrator/core"
)

type Stats struct {
	Schedule core.Schedule
	Schema   core.Schema
}

func (s *Service) GetStats(ctx context.Context) (Stats, error) {
	schedule, err := s.store.LoadSchedule(ctx)
	if err != nil {
		return Stats{}, err
	}

	schema, err := s.store.LoadSchema(ctx)
	if err != nil {
		return Stats{}, err
	}

	return Stats{
		Schedule: schedule,
		Schema:   schema,
	}, nil
}
