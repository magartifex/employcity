package api

import (
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/magartifex/employcity/constants"
	v1 "github.com/magartifex/employcity/proto/v1"
)

// Api server
type Api struct {
	domain domain

	v1.UnimplementedCRUDServer
}

func New(domain domain) *Api {
	return &Api{
		domain: domain,
	}
}

func (a *Api) Connect(grpcServer *grpc.Server) {
	v1.RegisterCRUDServer(grpcServer, a)
}

func parseError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, constants.ErrKeyNotExists) {
		return status.Error(codes.NotFound, "")
	}

	if errors.Is(err, constants.ErrInvalidDataTypeForKey) {
		return status.Error(codes.InvalidArgument, "key")
	}

	if errors.Is(err, constants.ErrInvalidDataTypeForValue) {
		return status.Error(codes.InvalidArgument, "value")
	}

	return status.Error(codes.Internal, "")
}
