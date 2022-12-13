package api

import (
	"context"

	v1 "github.com/magartifex/employcity/proto/v1"
)

func (a *Api) GetByKey(ctx context.Context, req *v1.GetByKey_Request) (*v1.GetByKey_Response, error) {
	entity, err := a.domain.GetByKey(ctx, req.GetKey())
	if err != nil {
		return nil, parseError(err)
	}

	return &v1.GetByKey_Response{Data: &v1.Data{
		Key:   entity.Key,
		Value: entity.Value,
	}}, nil
}
