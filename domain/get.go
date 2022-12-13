package domain

import (
	"context"
	"fmt"

	"github.com/magartifex/employcity/entity/data"
)

func (d *Domain) Get(ctx context.Context) ([]*data.Entity, error) {
	entities, err := d.store.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("get list error: %w", err)
	}

	return entities, nil
}
