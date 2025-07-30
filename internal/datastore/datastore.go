package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

func NewClient(ctx context.Context, projectID string) (*datastore.Client, error) {
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("datastore client error: %w", err)
	}
	return client, nil
}
