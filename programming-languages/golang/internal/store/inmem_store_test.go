package store

import (
	"fmt"
	"testing"
)

func TestStoreGet(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		wanted []byte
		err    error
	}{
		{
			name:   "can get existing key",
			key:    "greeting",
			wanted: []byte("hola, mundo"),
			err:    nil,
		},
		{
			name:   "errors when get non existing key",
			key:    "goodbye",
			wanted: nil,
			err:    fmt.Errorf(""),
		},
	}

	store := InMemoryStore{
		"greeting": []byte("hola, mundo"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := store.Get(tc.key)
			if err != nil && tc.err == nil {
				t.Errorf("got unexpected error %e", err)
			}

			if string(tc.wanted) != string(got) {
				t.Errorf("wanted %s, got %s", string(tc.wanted), string(got))
			}
		})
	}
}

func TestStoreSet(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value []byte
		err   error
	}{
		{
			name:  "can set new key",
			key:   "greeting",
			value: []byte("hola, mundo"),
			err:   nil,
		},
		{
			name:  "can set existing key",
			key:   "goodbye",
			value: []byte("adios, mundo"),
			err:   nil,
		},
	}

	store := InMemoryStore{
		"goodbye": []byte("chao, mundo"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := store.Set(tc.key, tc.value)
			if err != nil && tc.err == nil {
				t.Errorf("got unexpected error %e", err)
			}

			got, err := store.Get(tc.key)
			if err != nil {
				t.Errorf("got unexpected error %e", err)
			}

			if string(tc.value) != string(got) {
				t.Errorf("wanted %s, got %s", string(tc.value), string(got))
			}
		})
	}
}

func TestStoreDelete(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		err  error
	}{
		{
			name: "can delete existing key",
			key:  "greeting",
			err:  nil,
		},
		{
			name: "can delete non-existing key",
			key:  "greeting",
			err:  nil,
		},
	}

	store := InMemoryStore{
		"greeting": []byte("hola, mundo"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := store.Delete(tc.key)
			if err != nil && tc.err == nil {
				t.Errorf("got unexpected error %e", err)
			}

			value, _ := store.Get(tc.key)
			if value != nil {
				t.Errorf("should've thrown")
			}
		})
	}
}

func TestStoreKeys(t *testing.T) {
	testCases := []struct {
		name string
		keys []string
		err  error
	}{
		{
			name: "can get all keys set",
			keys: []string{"a", "b"},
			err:  nil,
		},
		{
			name: "can get if no keys set",
			keys: []string{},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := InMemoryStore{}
			for _, key := range tc.keys {
				store.Set(key, []byte(key))
			}

			keys, _ := store.Keys()
			if len(keys) != len(tc.keys) {
				t.Errorf("wanted length %d, got %d", len(tc.keys), len(keys))
			}
		})
	}
}
