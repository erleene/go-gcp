package main

import (
	"log"

	store "github.com/erleene/go-gcp/storage"
)

func main() {
	projectID := "name-of-project"

	fold, err := store.New(projectID)
	if err != nil {
		log.Fatal(err)
	}

	if err = fold.Delete("name-of-bucket", 14); err != nil {
		log.Fatal(err)
	}
}
