package main

import (
	"log"

	store "github.com/erleene/go-gcp/storage"
)

func main() {
	projectID := "beamery-preview"

	fold, err := store.New(projectID)
	if err != nil {
		log.Fatal(err)
	}

	if err = fold.Delete("uatscreenshots", 14); err != nil {
		log.Fatal(err)
	}
}
