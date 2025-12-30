package yaml

import (
	"context"
	"io/fs"
	"os"

	"github.com/jpuriol/cuadrator/core"
	"gopkg.in/yaml.v3"
)

type Store struct {
	fs               fs.FS
	participantsFile string
	schemaFile       string
	scheduleFile     string
}

func New(participants, schema, schedule string) *Store {
	return &Store{
		fs:               os.DirFS("."),
		participantsFile: participants,
		schemaFile:       schema,
		scheduleFile:     schedule,
	}
}

func (s *Store) LoadParticipants(ctx context.Context) (core.Participants, error) {
	data, err := fs.ReadFile(s.fs, s.participantsFile)
	if err != nil {
		return nil, err
	}

	var names []string
	if err := yaml.Unmarshal(data, &names); err != nil {
		return nil, err
	}

	participants := make(core.Participants)
	for _, name := range names {
		participants[name] = struct{}{}
	}
	return participants, nil
}

func (s *Store) LoadSchema(ctx context.Context) (core.Schema, error) {
	data, err := fs.ReadFile(s.fs, s.schemaFile)
	if err != nil {
		return core.Schema{}, err
	}

	var schema core.Schema
	if err := yaml.Unmarshal(data, &schema); err != nil {
		return core.Schema{}, err
	}
	return schema, nil
}

func (s *Store) LoadSchedule(ctx context.Context) (core.Schedule, error) {
	data, err := fs.ReadFile(s.fs, s.scheduleFile)
	if err != nil {
		return nil, err
	}

	var schedule core.Schedule
	if err := yaml.Unmarshal(data, &schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}
