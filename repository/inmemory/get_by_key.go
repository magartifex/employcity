package inmemory

import (
	"context"

	"github.com/magartifex/employcity/constants"
	"github.com/magartifex/employcity/entity/data"
)

func (s *Store) GetByKey(_ context.Context, key string) (*data.Entity, error) {
	res, ok := s.storage.Load(key)
	if !ok {
		return nil, constants.ErrKeyNotExists
	}

	value, ok := res.(string)
	if !ok {
		return nil, constants.ErrInvalidDataTypeForValue
	}

	return &data.Entity{
		Key:   key,
		Value: value,
	}, nil
}
