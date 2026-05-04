package main

import (
	"learning-golang/cmd"
	"learning-golang/internal/store"
)

func main() {
	store := store.InMemoryStore{}
	cmd.Execute(store)
}
