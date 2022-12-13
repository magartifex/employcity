package domain

import (
	"context"
	"fmt"
)

func (d *Domain) Delete(ctx context.Context, key string) error {
	err := d.store.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("delete key: %s, error: %w", key, err)
	}

	return nil
}
