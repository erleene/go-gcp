package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	store "github.com/erleene/go-gcp/storage"
)

type Contents struct {
	Name        string
	TimeCreated time.Time
}

func main() {
	projectID := "beamery-preview"

	//projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	// if projectID == "" {
	// 	fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
	// 	os.Exit(1)
	// }
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//list all buckets
	storageBucket, err := store.ListBucket(client, projectID)
	if err != nil {
		log.Fatal(err)
	}

	for _, bucketName := range storageBucket {
		if bucketName == "uatscreenshots" {
			fmt.Println("UAT Screenshot bucket exist")

			err := store.DeleteContentsOfBucket(client, projectID, bucketName)
			if err != nil {
				log.Fatal(err)
			}

			// //listcontents
			// contents, err := store.ListContentsOfBucket(client, projectID, bucketName)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// fmt.Printf("Contents of %s:\n %v \n", bucketName, contents)
			// //loop through contents to delete
		}
	}
}
