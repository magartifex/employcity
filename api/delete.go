package api

import (
	"context"

	v1 "github.com/magartifex/employcity/proto/v1"
)

func (a *Api) Delete(ctx context.Context, req *v1.Delete_Request) (*v1.Empty, error) {
	return &v1.Empty{}, parseError(a.domain.Delete(ctx, req.GetKey()))
}
