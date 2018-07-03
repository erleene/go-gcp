package storage

import (
	"context"

	"cloud.google.com/go/storage"
)

ctx := context.Background()
client, err := storage.NewClient(ctx)
if err != nil {
	return err
}
