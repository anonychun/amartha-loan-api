package entity

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Attachment struct {
	Id         uuid.UUID
	ObjectName string
	FileName   string
	ByteSize   int64
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Content []byte `gorm:"-"`
}

func (a *Attachment) BeforeCreate(tx *gorm.DB) error {
	a.Id = uuid.Must(uuid.NewV7())
	return nil
}

func NewAttachmentFromFile(file *os.File) (*Attachment, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &Attachment{
		ObjectName: ulid.Make().String() + filepath.Ext(fileInfo.Name()),
		FileName:   fileInfo.Name(),
		ByteSize:   fileInfo.Size(),

		Content: content,
	}, nil
}

func NewAttachmentFromMultipartFileHeader(fileHeader *multipart.FileHeader) (*Attachment, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &Attachment{
		ObjectName: ulid.Make().String() + filepath.Ext(fileHeader.Filename),
		FileName:   fileHeader.Filename,
		ByteSize:   fileHeader.Size,

		Content: content,
	}, nil
}
