package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

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
	storageBuckets, err := listBuckets(client, projectID)
	if err != nil {
		log.Fatal(err)
	}

	//print the buckets
	fmt.Println(storageBuckets)

	for _, bucketName := range storageBuckets {
		if bucketName == "uatscreenshots" {
			fmt.Println("UAT Screenshot bucket exist")

			// //list all the contents of the bucket
			// query := &storage.Query{}

		}
	}

}

func listBuckets(client *storage.Client, projectID string) ([]string, error) {
	ctx := context.Background()

	var buckets []string

	it := client.Buckets(ctx, projectID)

	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
	}
	return buckets, nil
}

func listContentsOfBucket(client *storage.Client, projectID string) ([]string, error) {
	ctx := context.Background()

	var contents []strings

	it := client.Bucket(projectID).Objects(ctx, nil)

	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		bucketObjects = append(contents)
	}

}
