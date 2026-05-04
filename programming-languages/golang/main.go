package main

import "fmt"

func main() {
	fmt.Println("project running")

	store := StoreImpl{}

	fmt.Println("initial keys: ", store.Keys())

	store.Set("greeting", []byte("hola, mundo"))
	store.Set("goodbye", []byte("adios, mundo"))

	fmt.Println("mid keys: ", store.Keys())
	greeting, err := store.Get("greeting")
	if err != nil {
		panic(fmt.Errorf("something went wrong: %e", err))
	}
	fmt.Println("greeting: ", string(greeting))

	_, err = store.Get("doesnt exist")
	if err == nil {
		panic(fmt.Errorf("shouldve errored"))
	}

	store.Delete("greeting")
	fmt.Println("final keys: ", store.Keys())
}
