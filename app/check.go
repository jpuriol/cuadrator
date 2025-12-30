package app

import (
	"context"
	"github.com/jpuriol/cuadrator/core"
)

func (s *Service) Check(ctx context.Context) error {
	schedule, err := s.store.LoadSchedule(ctx)
	if err != nil {
		return err
	}

	participants, err := s.store.LoadParticipants(ctx)
	if err != nil {
		return err
	}

	return core.ValidateSchedule(schedule, participants)
}
