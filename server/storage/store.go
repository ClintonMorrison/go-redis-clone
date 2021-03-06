package storage

import (
	"fmt"
	"go-redis-clone/server/primitives"
)

type Record struct {
	Value primitives.Value
}

type Store struct {
	Records map[string]*Record
}

func NewStore() *Store {
	store := Store{}
	store.Records = make(map[string]*Record, 0)

	return &store
}

func (s *Store) Set(key string, value primitives.Value) {
	s.Records[key] = &Record{value}
}

func (s *Store) Get(key string) primitives.Value {
	return s.Records[key].Value
}

func (s *Store) Delete(key string) primitives.Value {
	value := s.Get(key)
	delete(s.Records, key)
	return value
}

func (s *Store) DeleteAll() {
	s.Records = make(map[string]*Record, 0)
}

func (s *Store) Keys() []string {
	keys := make([]string, len(s.Records))
	i := 0
	for key := range s.Records {
		keys[i] = key
		i++
	}

	return keys
}



func (s *Store) Print() {
	for key, record := range s.Records {
		fmt.Println(key, "-->", record.Value)
	}
}