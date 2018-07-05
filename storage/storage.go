package storage

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// //Contents...
type Contents struct {
	Name        string
	TimeCreated time.Time
}

//List Bucket UAT Screenshot
func ListBucket(client *storage.Client, projectID string) ([]string, error) {
	ctx := context.Background()

	var bucket []string

	it := client.Buckets(ctx, projectID)

	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		//check if UAT screenshot exist
		switch {
		case battrs.Name == "uatscreenshots":
			bucket = append(bucket, battrs.Name)
		default:
			break
		}
	}
	return bucket, nil
}

//DeleteContentsOfBucket
func DeleteContentsOfBucket(client *storage.Client, projectID string, bucketName string) error {
	ctx := context.Background()
	now := time.Now()
	thresh := 60

	it := client.Bucket(bucketName).Objects(ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil && err != iterator.Done {
			return err
		}

		//check if it is 45 days older
		diff := now.Sub(attrs.Created)
		days := int(diff.Hours() / 24)

		if days >= thresh {
			//add to a list of contents to delete
			fmt.Println(attrs.Name)
			fmt.Println(attrs.Created.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
			fmt.Println(days)
			//
			fmt.Printf("Deleting...%s", attrs.Name)
			// if err := client.Bucket(bucketName).Object(attrs.Name).Delete(ctx); err != nil {
			// 	log.Fatal(err)
			//	}
		}
	}
	return nil
}
