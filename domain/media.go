package domain

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type ImageService interface {
	// Reads the image data from a reader and saves it and then returns image object.
	Save(ctx context.Context, form ImageUploadInput) (*Image, error)

	// Delete an image.
	Delete(ctx context.Context, id uuid.UUID) error

	// Get an image.
	One(ctx context.Context, id uuid.UUID) (*Image, error)

	// Get all images in a set of ids.
	AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*Image, error)
}

type Image struct {
	ID        uuid.UUID   `json:"id"`
	Title     null.String `json:"title"`
	Link      string      `json:"link"`
	FileHash  []byte      `json:"-"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type VideoService interface {
	// Reads the video data from reader, saves it to disk and returns video object.
	Save(ctx context.Context, input VideoUploadInput) (*Video, error)

	// Update the video.
	Update(ctx context.Context, input VideoUpdateInput) error

	// Delete a video.
	Delete(ctx context.Context, id uuid.UUID) error

	// Get Videos by ids.
	AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*Video, error)
}

type Video struct {
	ID          uuid.UUID   `json:"id"`
	Title       null.String `json:"title"`
	FileHash    string      `json:"fileHash"`
	ThumbnailID uuid.UUID   `json:"thumbnailId"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

//------ MEDIA FORMS -----------

type ImageUploadInput struct {
	Title       null.String
	File        io.ReadSeeker
	ContentType string
}

func (i ImageUploadInput) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.ContentType, vd.Required, vd.By(func(_ interface{}) error {
			splitContentType := strings.Split(i.ContentType, "/")

			if splitContentType[0] != "image" {
				return errors.New("invalid file type, the uploaded file is not an image")
			}

			return nil
		})),
	)
}

type ImageUpdateInput struct {
	ID          uuid.UUID
	Title       null.String
	File        io.ReadSeeker
	ContentType string
}

func (i ImageUpdateInput) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.ID, is.UUIDv4),
		vd.Field(&i.ContentType, vd.Required, vd.By(func(_ interface{}) error {
			splitContentType := strings.Split(i.ContentType, "/")

			if splitContentType[0] != "image" {
				return errors.New("invalid file type, the uploaded file is not an image")
			}

			return nil
		})),
	)
}

type VideoUploadInput struct {
	Title             null.String
	File              io.Reader
	CustomThumbnailID uuid.NullUUID
}

func (v VideoUploadInput) Validate() error {
	return vd.ValidateStruct(&v,
		vd.Field(&v.File, vd.By(func(_ interface{}) error {
			buf := make([]byte, 512)
			nBytes, err := v.File.Read(buf)
			if err != nil {
				return err
			}
			log.Printf("validation read %d bytes from reader\n", nBytes)

			mimeType := http.DetectContentType(buf)

			splitMime := strings.Split(mimeType, "/")

			if splitMime[0] != "video" {
				return errors.New("invalid file type, the uploaded file is not an video")
			}

			return nil
		})),
	)
}

type VideoUpdateInput struct {
	ID                uuid.UUID
	Title             null.String
	File              io.Reader
	CustomThumbnailID uuid.NullUUID
}

func (v VideoUpdateInput) Validate() error {
	return vd.ValidateStruct(&v,
		vd.Field(&v.File, is.UUIDv4),
		vd.Field(&v.File, vd.By(func(_ interface{}) error {
			buf := make([]byte, 512)
			nBytes, err := v.File.Read(buf)
			if err != nil {
				return err
			}
			log.Printf("validation read %d bytes from reader\n", nBytes)

			mimeType := http.DetectContentType(buf)

			splitMime := strings.Split(mimeType, "/")

			if splitMime[0] != "video" {
				return errors.New("invalid file type, the uploaded file is not an video")
			}

			return nil
		})),
	)
}
