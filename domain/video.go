package domain

import (
	"context"
	"errors"
	"time"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/input"
)

type VideoService interface {
	// Reads the video data from reader, saves it to disk and returns video object.
	Save(ctx context.Context, input input.VideoUpload) (*Video, error)

	// Update the video.
	Update(ctx context.Context, input input.VideoUpdate) error

	// Delete a video.
	Delete(ctx context.Context, id uuid.UUID) error

	// Get Videos.
	Videos(ctx context.Context, filter input.MediaFilter) ([]Video, error)
}

type VideoRepository interface {
	SaveOrUpdate(ctx context.Context, video *Video) error
	Delete(ctx context.Context, id uuid.UUID) error
	OneByID(ctx context.Context, id uuid.UUID) (*Video, error)
	All(ctx context.Context, ids ...uuid.UUID) ([]Video, error)
	Many(ctx context.Context, filter input.MediaFilter) ([]Video, error)
	CheckHashExists(ctx context.Context, hash string) (bool, error)
}

type Video struct {
	ID          uuid.UUID   `json:"id"`
	Title       null.String `json:"title"`
	FileHash    string      `json:"fileHash"`
	ThumbnailID uuid.UUID   `json:"thumbnailId"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func (v Video) Validate() error {
	return vd.ValidateStruct(&v,
		vd.Field(&v.ID, vd.By(func(_ interface{}) error {
			if v.ID == uuid.Nil {
				return errors.New("id cannot be blank")
			}
			return nil
		})),
		vd.Field(&v.FileHash, vd.Required),
		vd.Field(&v.CreatedAt, vd.By(func(_ interface{}) error {
			if v.CreatedAt.IsZero() {
				return errors.New("invalid createdAt time")
			}
			return nil
		})),
		vd.Field(&v.UpdatedAt, vd.By(func(_ interface{}) error {
			if v.UpdatedAt.IsZero() {
				return errors.New("invalid updatedAt time")
			}
			return nil
		})),
	)
}

func (v Video) IsEqual(c *Video) bool {
	return v.ID == c.ID &&
		v.Title == c.Title &&
		v.FileHash == c.FileHash &&
		v.CreatedAt.Equal(c.CreatedAt) &&
		v.UpdatedAt.Equal(c.UpdatedAt)
}
