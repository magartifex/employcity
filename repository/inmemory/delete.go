package inmemory

import "context"

func (s *Store) Delete(_ context.Context, key string) error {
	s.storage.Delete(key)
	return nil
}
