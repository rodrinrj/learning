package main

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

	store := StoreImpl{
		"greeting": []byte("hola, mundo"),
	}

	for _, tc := range testCases {
		got, err := store.Get(tc.key)
		if err != nil && tc.err == nil {
			t.Errorf("%s: got unexpected error %e", tc.name, err)
		}

		if string(tc.wanted) != string(got) {
			t.Errorf("%s: wanted %s, got %s", tc.name, string(tc.wanted), string(got))
		}
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

	store := StoreImpl{
		"goodbye": []byte("chao, mundo"),
	}

	for _, tc := range testCases {
		err := store.Set(tc.key, tc.value)
		if err != nil && tc.err == nil {
			t.Errorf("%s: got unexpected error %e", tc.name, err)
		}

		got, err := store.Get(tc.key)
		if err != nil {
			t.Errorf("%s: got unexpected error %e", tc.name, err)
		}

		if string(tc.value) != string(got) {
			t.Errorf("%s: wanted %s, got %s", tc.name, string(tc.value), string(got))
		}
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

	store := StoreImpl{
		"greeting": []byte("hola, mundo"),
	}

	for _, tc := range testCases {
		err := store.Delete(tc.key)
		if err != nil && tc.err == nil {
			t.Errorf("%s: got unexpected error %e", tc.name, err)
		}

		value, _ := store.Get(tc.key)
		if value != nil {
			t.Errorf("%s: should've thrown", tc.name)
		}
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
		store := StoreImpl{}
		for _, key := range tc.keys {
			store.Set(key, []byte(key))
		}

		keys := store.Keys()
		if len(keys) != len(tc.keys) {
			t.Errorf("%s: wanted length of %d got %d", tc.name, len(tc.keys), len(keys))
		}
	}
}
