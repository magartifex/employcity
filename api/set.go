package api

import (
	"context"

	v1 "github.com/magartifex/employcity/proto/v1"
)

func (a *Api) Set(ctx context.Context, req *v1.Set_Request) (*v1.Set_Response, error) {
	entity, err := a.domain.Set(ctx, req.GetValue())
	if err != nil {
		return nil, parseError(err)
	}

	return &v1.Set_Response{Data: &v1.Data{
		Key:   entity.Key,
		Value: entity.Value,
	}}, nil
}
