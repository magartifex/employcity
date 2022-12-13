package config

import "github.com/kelseyhightower/envconfig"

// Params config
type Params struct {
	GrpcPort      int  `envconfig:"GRPC_PORT" required:"true"`
	IsInMemory    bool `envconfig:"IS_IN_MEMORY" required:"true"`
	MemcachedPort int  `envconfig:"MEMCACHED_PORT"`
}

// New returns actual config instance
func New() (*Params, error) {
	cfg := &Params{}
	return cfg, envconfig.Process("", cfg)
}
