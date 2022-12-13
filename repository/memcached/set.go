package memcached

import (
	"context"

	"github.com/magartifex/employcity/entity/data"
	v1 "github.com/magartifex/employcity/proto/v1"
)

func (s *Store) Set(ctx context.Context, value string) (*data.Entity, error) {
	resp, err := s.client.Set(ctx, &v1.Set_Request{Value: value})
	if err != nil {
		return nil, parseError(err)
	}

	return &data.Entity{
		Key:   resp.GetData().GetKey(),
		Value: resp.GetData().GetValue(),
	}, nil
}
