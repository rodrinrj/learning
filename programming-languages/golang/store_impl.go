package main

import "fmt"

type StoreImpl map[string][]byte

// Gets the value if the key exists, returns error if it does not
func (s StoreImpl) Get(key string) ([]byte, error) {
	value, exists := s[key]
	if !exists {
		return nil, fmt.Errorf("key does not exist")
	}

	return value, nil
}

func (s StoreImpl) Set(key string, value []byte) error {
	s[key] = value
	return nil
}

func (s StoreImpl) Delete(key string) error {
	delete(s, key)
	return nil
}

func (s StoreImpl) Keys() []string {
	keys := []string{}
	for key := range s {
		keys = append(keys, key)
	}

	return keys
}
