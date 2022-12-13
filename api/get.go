package api

import (
	"context"

	v1 "github.com/magartifex/employcity/proto/v1"
)

func (a *Api) Get(ctx context.Context, _ *v1.Empty) (*v1.Get_Response, error) {
	entities, err := a.domain.Get(ctx)
	if err != nil {
		return nil, parseError(err)
	}

	list := make([]*v1.Data, len(entities))
	for i, entity := range entities {
		list[i] = &v1.Data{
			Key:   entity.Key,
			Value: entity.Value,
		}
	}

	return &v1.Get_Response{List: list}, nil
}
