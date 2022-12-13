package inmemory

import (
	"context"

	"github.com/magartifex/employcity/constants"
	"github.com/magartifex/employcity/entity/data"
)

func (s *Store) Get(ctx context.Context) ([]*data.Entity, error) {
	list := make([]*data.Entity, 0)
	var err error

	s.storage.Range(func(key any, value any) bool {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return false

		default:
			k, ok := key.(string)
			if !ok {
				err = constants.ErrInvalidDataTypeForKey
				return false
			}

			v, ok := value.(string)
			if !ok {
				err = constants.ErrInvalidDataTypeForValue
				return false
			}

			list = append(list, &data.Entity{
				Key:   k,
				Value: v,
			})

			return true
		}
	})

	return list, err
}
