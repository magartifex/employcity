package memcached

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/magartifex/employcity/constants"
	v1 "github.com/magartifex/employcity/proto/v1"
)

func (s *Store) Delete(ctx context.Context, key string) error {
	_, err := s.client.Delete(ctx, &v1.Delete_Request{Key: key})
	if err != nil {
		return parseError(err)
	}

	return nil
}

func parseError(err error) error {
	if err == nil {
		return nil
	}

	if status.Code(err) == codes.NotFound {
		return constants.ErrKeyNotExists
	}

	if status.Code(err) == codes.InvalidArgument {
		switch status.Convert(err).Message() {
		case "key":
			return constants.ErrInvalidDataTypeForKey

		case "value":
			return constants.ErrInvalidDataTypeForKey

		default:
			return err
		}
	}

	return err
}
