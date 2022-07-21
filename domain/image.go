package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/thanishsid/goserver/internal/input"
)

type ImageService interface {
	// Reads the image data from a reader and saves it and then returns image object.
	Save(ctx context.Context, input input.ImageUpload) (*Image, error)

	// Update the image.
	Update(ctx context.Context, input input.ImageUpdate) error

	// Delete an image.
	Delete(ctx context.Context, id uuid.UUID) error
}

type ImageRepository interface {
	Save(ctx context.Context, image *Image) error
	Update(ctx context.Context, image *Image) error
	Delete(ctx context.Context, id uuid.UUID) error
	OneByID(ctx context.Context, id uuid.UUID) (*Image, error)
	All(ctx context.Context, ids ...uuid.UUID) ([]Image, error)
}

type Image struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	FileHash  string    `json:"fileHash"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
