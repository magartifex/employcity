package api

import (
	"context"

	"github.com/magartifex/employcity/entity/data"
)

type domain interface {
	Set(context.Context, string) (*data.Entity, error)
	Get(context.Context) ([]*data.Entity, error)
	GetByKey(context.Context, string) (*data.Entity, error)
	Delete(context.Context, string) error
}
