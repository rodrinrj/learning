package store

import (
	"maps"
	"slices"
)

type InMemoryStore map[string][]byte

// Gets the value if the key exists, returns error if it does not
func (s InMemoryStore) Get(key string) ([]byte, error) {
	value, exists := s[key]
	if !exists {
		return nil, ErrKeyNotFound
	}

	return value, nil
}

func (s InMemoryStore) Set(key string, value []byte) error {
	if s == nil {
		return ErrSetFailed
	}

	s[key] = value
	return nil
}

func (s InMemoryStore) Delete(key string) error {
	if s == nil {
		return ErrDelFailed
	}

	delete(s, key)
	return nil
}

func (s InMemoryStore) Keys() ([]string, error) {
	if s == nil {
		return nil, ErrKeysFailed
	}

	keys := slices.Collect(maps.Keys(s))

	return keys, nil
}
