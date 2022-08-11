package domain

import (
	"context"
	"errors"
	"strings"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type ImageService interface {
	// Reads the image data from a reader and saves it and then returns image object.
	Save(ctx context.Context, input ImageUploadInput) (*Image, error)

	// Delete an image.
	Delete(ctx context.Context, id uuid.UUID) error

	// Get an image.
	One(ctx context.Context, id uuid.UUID) (*Image, error)

	// Get all images in a set of ids.
	AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*Image, error)
}

type Image struct {
	ID        uuid.UUID `json:"id"`
	FileName  string    `json:"-"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VideoService interface {
	// Reads the video data from reader, saves it to disk and returns video object.
	Save(ctx context.Context, input VideoUploadInput) (*Video, error)

	// Delete a video.
	Delete(ctx context.Context, id uuid.UUID) error

	// Get a video.
	One(ctx context.Context, id uuid.UUID) (*Video, error)

	// Get Videos by ids.
	AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*Video, error)
}

type Video struct {
	ID          uuid.UUID `json:"id"`
	FileName    string    `json:"-"`
	Link        string    `json:"link"`
	ThumbnailID uuid.UUID `json:"thumbnailId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

//------ MEDIA FORMS -----------

type ImageUploadInput struct {
	FileUploadData
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

type VideoUploadInput struct {
	FileUploadData
	CustomThumbnailID uuid.NullUUID
}

func (v VideoUploadInput) Validate() error {
	return vd.ValidateStruct(&v,
		vd.Field(&v.ContentType, vd.Required, vd.By(func(value interface{}) error {
			splitContentType := strings.Split(v.ContentType, "/")

			if splitContentType[0] != "video" {
				return errors.New("invalid file type, the uploaded file is not a video")
			}

			return nil
		})),
	)
}
