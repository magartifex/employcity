package inmemory

import (
	"context"

	"github.com/google/uuid"

	"github.com/magartifex/employcity/entity/data"
)

func (s *Store) Set(_ context.Context, value string) (*data.Entity, error) {
	key := uuid.New().String()
	s.storage.Store(key, value)
	return &data.Entity{
		Key:   key,
		Value: value,
	}, nil
}
