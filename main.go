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
	storageBuckets, err := store.ListBuckets(client, projectID)
	if err != nil {
		log.Fatal(err)
	}

	for _, bucketName := range storageBuckets {
		if bucketName == "uatscreenshots" {
			fmt.Println("UAT Screenshot bucket exist")

			//listcontents
			contents, err := store.ListContentsOfBucket(client, projectID, bucketName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Contents of %s:\n %v \n", bucketName, contents)
			//loop through array struct
			//change to type map[type]type
		}
	}
}
