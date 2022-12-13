package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Conn(logger *log.Logger, port int, callbacks ...interface{ Connect(*grpc.Server) }) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("net.Listen() error = %w", err)
	}

	defer func() {
		err := listener.Close()
		if err != nil {
			logger.Printf("listener.Close() error = %s\n", err)
			return
		}
	}()

	grpcServer := grpc.NewServer()

	for _, callback := range callbacks {
		callback.Connect(grpcServer)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("grpcServer.Serve() error = %w", err)
	}

	return nil
}
