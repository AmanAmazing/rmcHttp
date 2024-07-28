package main

import (
	"fmt"
	"sync"
)

// TODO: Add sharding to improve performance

type Store struct {
	data map[string]any
	sync.RWMutex
}

func NewStore() Store {
	return Store{data: make(map[string]any)}
}

func (s *Store) Clear() {
	clear(s.data)
	return
}

func (s *Store) Get(k string) (any, bool) {
	s.RLock()
	val := s.data[k]
	s.RUnlock()
	return val, val != nil
}

func (s *Store) Set(k string, v any) {
	s.Lock()
	s.data[k] = v
	s.Unlock()
}
func (s *Store) Delete(k string) {
	s.Lock()
	delete(s.data, k)
	s.Unlock()
}

func (s *Store) Contains(k string) bool {
	s.RLock()
	val := s.data[k]
	s.RUnlock()
	return val != nil
}
func (s *Store) Keys() []string {
	s.RLock()
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	s.RUnlock()
	return keys
}

func (s *Store) PrintAll() {
	s.RLock()
	for key, value := range s.data {
		fmt.Println("Key:", key, "value:", value)
	}
	s.RUnlock()
}

func main() {
	cache := NewStore()
	cache.Set("charile", "dog")
	cache.Set("brazy", "wolf")
	cache.Set("piggy", "pig")
	cache.PrintAll()
}
