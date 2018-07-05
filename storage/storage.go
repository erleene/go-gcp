package storage

import (
	"context"
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

	//var contents []string //to store the contents of a bucket
	var contents []Contents
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
		//add each one content of bucket to struct array
		contents = append(contents, Contents{
			Name:        attrs.Name,
			TimeCreated: attrs.Created,
		})

		for _, v := range contents {

			now := time.Now()
			diff := now.Sub(v.TimeCreated)
			//convert diff to days instead of hours
			days := int(diff.Hours() / 24)
			//print those that are 45 days old
			thresh := 45
			if days >= thresh {
				//fmt.Println(v.TimeCreated.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
				//fmt.Println(days)
				//add to a list of contents to delete
				olderContents = append(olderContents, Contents{
					Name:        v.Name,
					TimeCreated: v.TimeCreated,
				})
			}
		}
	}
	return olderContents, nil
}
