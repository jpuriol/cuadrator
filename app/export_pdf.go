package app

import (
	"context"
	"io"
)

func (s *Service) ExportPDF(ctx context.Context, w io.Writer) error {
	schedule, err := s.store.LoadSchedule(ctx)
	if err != nil {
		return err
	}

	participants, err := s.store.LoadParticipants(ctx)
	if err != nil {
		return err
	}

	schema, err := s.store.LoadSchema(ctx)
	if err != nil {
		return err
	}

	return s.generator.Generate(ctx, schedule, participants, schema, w)
}
