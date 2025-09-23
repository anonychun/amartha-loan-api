package dto

import (
	"context"

	"github.com/anonychun/amartha-loan-api/internal/entity"
	"github.com/anonychun/amartha-loan-api/internal/storage"
)

type Attachment struct {
	Id       string `json:"id"`
	FileName string `json:"fileName"`
	Url      string `json:"url"`
}

func NewAttachment(s3 *storage.S3, attachment *entity.Attachment) (*Attachment, error) {
	url, err := s3.PresignedUrl(context.Background(), attachment.ObjectName)
	if err != nil {
		return nil, err
	}

	return &Attachment{
		Id:       attachment.Id.String(),
		FileName: attachment.FileName,
		Url:      url,
	}, nil
}
