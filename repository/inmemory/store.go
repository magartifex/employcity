package inmemory

import "sync"

// Store in memory
type Store struct {
	storage sync.Map
}

// New store
func New() *Store {
	return &Store{
		storage: sync.Map{},
	}
}
