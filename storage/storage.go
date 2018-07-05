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

//List Buckets
func ListBuckets(client *storage.Client, projectID string) ([]string, error) {
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

//ListContentsOfBucket
func ListContentsOfBucket(client *storage.Client, projectID string, bucketName string) ([]Contents, error) {
	ctx := context.Background()

	var olderContents []Contents

	it := client.Bucket(bucketName).Objects(ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		//check if it is 45 days older
		now := time.Now()
		diff := now.Sub(attrs.Created)
		days := int(diff.Hours() / 24)
		thresh := 45

		if days >= thresh {
			//add to a list of contents to delete
			fmt.Println(attrs.Name)
			fmt.Println(attrs.Created.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
			fmt.Println(days)

			//add the contents that needs to be deleted
			olderContents = append(olderContents, Contents{
				Name:        attrs.Name,
				TimeCreated: attrs.Created,
			})
		}
	}
	return olderContents, nil
}
