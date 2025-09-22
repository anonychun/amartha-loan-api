package storage

import (
	"bytes"
	"context"

	"github.com/anonychun/amartha-loan-api/internal/bootstrap"
	"github.com/anonychun/amartha-loan-api/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewS3)
}

type S3 struct {
	client *minio.Client
	config *config.Config
}

func NewS3(i do.Injector) (*S3, error) {
	cfg, err := do.Invoke[*config.Config](i)
	if err != nil {
		return nil, err
	}

	client, err := minio.New(cfg.Storage.S3.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Storage.S3.AccessKeyId, cfg.Storage.S3.SecretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return &S3{
		client: client,
		config: cfg,
	}, nil
}

func (s *S3) Put(ctx context.Context, objectName string, content []byte, byteSize int64) error {
	_, err := s.client.PutObject(
		ctx,
		s.config.Storage.S3.Bucket,
		objectName,
		bytes.NewReader(content),
		byteSize,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3) PresignedUrl(ctx context.Context, objectName string) (string, error) {
	url, err := s.client.PresignedGetObject(
		ctx,
		s.config.Storage.S3.Bucket,
		objectName,
		s.config.Storage.S3.UrlExpiration,
		nil,
	)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}
