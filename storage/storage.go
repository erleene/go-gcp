package storage

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// //Contents...
type Contents struct {
	Name        string
	TimeCreated time.Time
}

type storager struct {
	ctx       context.Context
	client    *storage.Client
	projectID string
}

// Folder is interfacing with the Google Cloud Bucket
type Folder interface {
	// Delete will remove all contents inside the bucket that are after the defined threshold
	Delete(bucketName string, threshold int) error
}

// New is to enable user to create a Folder instance
func New(projectID string) (Folder, error) {
	if projectID == "" {
		return &storager{}, errors.New("Please provide a project ID")
	}
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return &storager{}, err
	}
	return &storager{
		ctx:       context.Background(),
		client:    client,
		projectID: projectID,
	}, nil
}

func (store *storager) Delete(bucketName string, threshold int) error {
	now := time.Now()

	it := store.client.Bucket(bucketName).Objects(store.ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil && err != iterator.Done {
			return err
		}

		diff := now.Sub(attrs.Created)
		days := int(diff.Hours() / 24)

		if days >= threshold {
			if err := store.client.Bucket(bucketName).Object(attrs.Name).Delete(store.ctx); err != nil {
				return err
			}
		}
	}
	return nil
}
