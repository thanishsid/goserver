package domain

import (
	"context"
	"errors"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/internal/input"
)

type ImageService interface {
	// Reads the image data from a reader and saves it and then returns image object.
	Save(ctx context.Context, input input.ImageUpload) (*Image, error)

	// Update the image.
	Update(ctx context.Context, input input.ImageUpdate) error

	// Delete an image.
	Delete(ctx context.Context, id uuid.UUID) error

	// Get images.
	Images(ctx context.Context, filter input.MediaFilter) (*ListWithCursor[Image], error)
}

type ImageRepository interface {
	SaveOrUpdate(ctx context.Context, image *Image) error
	Delete(ctx context.Context, id uuid.UUID) error
	OneByID(ctx context.Context, id uuid.UUID) (*Image, error)
	All(ctx context.Context, ids ...uuid.UUID) ([]Image, error)
	Many(ctx context.Context, params input.MediaFilterBase) ([]Image, error)
	CheckHashExists(ctx context.Context, hash string) (bool, error)
}

type Image struct {
	ID        uuid.UUID   `json:"id"`
	Title     null.String `json:"title"`
	Link      string      `json:"link"`
	FileHash  string      `json:"-"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func (i Image) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.ID, is.UUIDv4),
		vd.Field(&i.Link, vd.Required),
		vd.Field(&i.FileHash, vd.Required),
		vd.Field(&i.CreatedAt, vd.By(func(_ interface{}) error {
			if i.CreatedAt.IsZero() {
				return errors.New("invalid createdAt time")
			}
			return nil
		})),
		vd.Field(&i.UpdatedAt, vd.By(func(_ interface{}) error {
			if i.UpdatedAt.IsZero() {
				return errors.New("invalid updatedAt time")
			}
			return nil
		})),
	)
}

// Check if the image object is equal to the provided image object.
func (i Image) IsEqual(c *Image) bool {
	return i.ID == c.ID &&
		i.Title == c.Title &&
		i.FileHash == c.FileHash &&
		i.CreatedAt == c.CreatedAt &&
		i.UpdatedAt == c.UpdatedAt
}

// Loads the image proxy link into the Image object.
func (i *Image) LoadImageProxyLink() {
	i.Link = config.C.ImageProxyLink
}
