type StorageConfig struct {
    Provider       string // "s3", "gcs", "minio"
    Endpoint       string // Empty for AWS/GCS, URL for MinIO
    Region         string
    ForcePathStyle bool   // true for MinIO
    Buckets        BucketConfig
}

type BucketConfig struct {
    Uploads string // e.g., "bastet-uploads"
    Exports string // e.g., "bastet-exports"
    Backups string // e.g., "bastet-backups"
}

func NewObjectStore(ctx context.Context, cfg StorageConfig) (ObjectStore, error) {
    switch cfg.Provider {
    case "s3", "minio":
        return NewS3Store(ctx, cfg.Endpoint, cfg.Region, cfg.ForcePathStyle)
    case "gcs":
        return NewGCSStore(ctx) // Implement similarly
    default:
        return nil, fmt.Errorf("unsupported storage provider: %s", cfg.Provider)
    }
}
