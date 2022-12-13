package domain

import (
	"context"
	"fmt"

	"github.com/magartifex/employcity/entity/data"
)

func (d *Domain) GetByKey(ctx context.Context, key string) (*data.Entity, error) {
	entity, err := d.store.GetByKey(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("get by key: %s, error: %w", key, err)
	}

	return entity, nil
}
