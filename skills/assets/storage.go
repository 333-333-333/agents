// internal/shared/storage/storage.go
package storage

import (
	"context"
	"io"
	"time"
)

type ObjectStore interface {
	Put(ctx context.Context, bucket, key string, reader io.Reader, opts PutOptions) error
	Get(ctx context.Context, bucket, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, bucket, key string) error
	Exists(ctx context.Context, bucket, key string) (bool, error)
	SignedURL(ctx context.Context, bucket, key string, expiry time.Duration) (string, error)
	List(ctx context.Context, bucket, prefix string) ([]ObjectInfo, error)
}

type PutOptions struct {
	ContentType string
	Metadata    map[string]string
}

type ObjectInfo struct {
	Key          string
	Size         int64
	LastModified time.Time
	ContentType  string
}
