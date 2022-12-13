package memcached

import (
	"context"

	"github.com/magartifex/employcity/entity/data"
	v1 "github.com/magartifex/employcity/proto/v1"
)

func (s *Store) Get(ctx context.Context) ([]*data.Entity, error) {
	resp, err := s.client.Get(ctx, &v1.Empty{})
	if err != nil {
		return nil, parseError(err)
	}

	list := make([]*data.Entity, len(resp.GetList()))

	for i, entity := range resp.GetList() {
		list[i] = &data.Entity{
			Key:   entity.GetKey(),
			Value: entity.GetValue(),
		}
	}

	return list, nil
}
