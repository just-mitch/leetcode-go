package main

import (
	"fmt"
	"log"

	"github.com/just-mitch/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Paul", "Mitch", "John"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(messages)
	}
}
