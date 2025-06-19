package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello there!")
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%+v\n",store)
	server := NewAPIServer(":3000", store)
	server.Run()

}
