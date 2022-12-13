package domain

import (
	"context"
	"fmt"

	"github.com/magartifex/employcity/entity/data"
)

func (d *Domain) Set(ctx context.Context, value string) (*data.Entity, error) {
	entity, err := d.store.Set(ctx, value)
	if err != nil {
		return nil, fmt.Errorf("error set value: %s, error: %w", value, err)
	}

	return entity, nil
}
