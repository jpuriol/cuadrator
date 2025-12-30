package core

import "errors"

var (
	ErrParticipantNotFound = errors.New("participant not found")
	ErrDuplicateAssignment = errors.New("duplicate assignment")
)
