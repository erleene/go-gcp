package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	storage "github.com/erleene/go-gcp/storage"
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
	storageBuckets, err := storage.ListBuckets(client, projectID)
	if err != nil {
		log.Fatal(err)
	}
	//print the buckets
	//fmt.Println(storageBuckets)

	for _, bucketName := range storageBuckets {
		if bucketName == "uatscreenshots" {
			fmt.Println("UAT Screenshot bucket exist")

			//listcontents
			contents, err := storage.ListContentsOfBucket(client, projectID, bucketName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Contents of %s:\n %v \n", bucketName, contents)
			//loop through array struct
			//change to type map[type]type

		}
	}
}

//
// //function to call to list all the buckets within a project
// func listBuckets(client *storage.Client, projectID string) ([]string, error) {
// 	ctx := context.Background()
//
// 	var buckets []string
//
// 	it := client.Buckets(ctx, projectID)
//
// 	for {
// 		battrs, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return nil, err
// 		}
// 		buckets = append(buckets, battrs.Name)
// 	}
// 	return buckets, nil
// }
//
// //list contents that are older than given days
// func listContentsOfBucket(client *storage.Client, projectID string, bucketName string) ([]Contents, error) {
// 	ctx := context.Background()
//
// 	//var contents []string //to store the contents of a bucket
// 	var contents []Contents
// 	var olderContents []Contents
//
// 	it := client.Bucket(bucketName).Objects(ctx, nil)
//
// 	for {
// 		attrs, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return nil, err
// 		}
// 		//add each one content of bucket to struct array
// 		contents = append(contents, Contents{
// 			Name:        attrs.Name,
// 			TimeCreated: attrs.Created,
// 		})
//
// 		for _, v := range contents {
//
// 			now := time.Now()
// 			diff := now.Sub(v.TimeCreated)
// 			//convert diff to days instead of hours
// 			days := int(diff.Hours() / 24)
// 			//print those that are 50 days old
// 			thresh := 45
// 			if days >= thresh {
// 				fmt.Println(v.TimeCreated.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
// 				fmt.Println(days)
// 				//add to a list of contents to delete
// 				olderContents = append(olderContents, Contents{
// 					Name:        attrs.Name,
// 					TimeCreated: attrs.Created,
// 				})
// 			}
// 		}
// 	}
// 	return olderContents, nil
// }
