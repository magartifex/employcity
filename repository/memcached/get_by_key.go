package memcached

import (
	"context"

	"github.com/magartifex/employcity/entity/data"
	v1 "github.com/magartifex/employcity/proto/v1"
)

func (s *Store) GetByKey(ctx context.Context, key string) (*data.Entity, error) {
	resp, err := s.client.GetByKey(ctx, &v1.GetByKey_Request{Key: key})
	if err != nil {
		return nil, parseError(err)
	}

	return &data.Entity{
		Key:   resp.GetData().GetKey(),
		Value: resp.GetData().GetValue(),
	}, nil
}
