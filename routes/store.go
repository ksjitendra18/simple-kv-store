// store.go

package routes

import (
	"sync"
	"time"
)

type Store struct {
	data map[string]valueWithExpiry
	mu   sync.RWMutex
}

type valueWithExpiry struct {
	value  string
	expiry time.Time
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]valueWithExpiry),
	}
}

func (s *Store) Set(key, value string, expiryTime time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var expiry time.Time
	if expiryTime > 0 {
		expiry = time.Now().Add(expiryTime)
	}

	s.data[key] = valueWithExpiry{
		value:  value,
		expiry: expiry,
	}

	go func() {
		for range time.Tick(time.Minute) {
			s.mu.Lock()
			for key, item := range s.data {
				if item.expiry.Before(time.Now()) {
					delete(s.data, key)
				}
			}
			s.mu.Unlock()
		}
	}()
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.data[key]
	if !ok || (item.expiry != time.Time{} && item.expiry.Before(time.Now())) {
		return "", false
	}

	return item.value, true
}
