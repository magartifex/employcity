package memcached

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	v1 "github.com/magartifex/employcity/proto/v1"
)

// Store in memory
type Store struct {
	client v1.CRUDClient
}

// New store
func New(port int) (*Store, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		return nil, fmt.Errorf("grpc.Dial() error = %w", err)
	}

	return &Store{
		client: v1.NewCRUDClient(conn),
	}, nil
}
