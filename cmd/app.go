package main

import (
	"log"

	"github.com/magartifex/employcity/api"
	"github.com/magartifex/employcity/config"
	"github.com/magartifex/employcity/domain"
	"github.com/magartifex/employcity/repository/inmemory"
	"github.com/magartifex/employcity/repository/memcached"
	"github.com/magartifex/employcity/server"
)

func main() {
	logger := log.Default()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config.New() error=%s", err)
		return
	}

	var store domain.Store

	if cfg.IsInMemory {
		store = inmemory.New()
	} else {
		store, err = memcached.New(cfg.MemcachedPort)
		if err != nil {
			log.Fatalf("memcached.New() error=%s", err)
			return
		}
	}

	dom := domain.New(store)
	apiServer := api.New(dom)

	err = server.Conn(logger, cfg.GrpcPort, apiServer)
	if err != nil {
		log.Fatalf("server.Conn() error=%s", err)
		return
	}
}
