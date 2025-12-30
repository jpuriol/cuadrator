package app

import (
	"context"
	"github.com/jpuriol/cuadrator/core"
	"io"
)

type Store interface {
	LoadParticipants(ctx context.Context) (core.Participants, error)
	LoadSchema(ctx context.Context) (core.Schema, error)
	LoadSchedule(ctx context.Context) (core.Schedule, error)
}

type PDFGenerator interface {
	Generate(ctx context.Context, schedule core.Schedule, participants core.Participants, schema core.Schema, w io.Writer) error
}

type Service struct {
	store     Store
	generator PDFGenerator
}

func NewService(store Store, generator PDFGenerator) *Service {
	return &Service{
		store:     store,
		generator: generator,
	}
}
