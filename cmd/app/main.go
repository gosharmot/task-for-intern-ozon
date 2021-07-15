package main

import (
	"log"
	"task-for-intern/internal/server"
)

func main() {
	log.SetFlags(log.Lshortfile)

	server := server.NewServer(":8080")

	if err := server.Listen(); err != nil {
		log.Fatal(err)
	}
}
