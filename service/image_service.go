package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/internal/input"
)

type imageService struct{}

var _ domain.ImageService = (*imageService)(nil)

// Reads the image data from a reader and saves it and then returns image object.
func (i *imageService) Save(ctx context.Context, input input.ImageUpload) (*domain.Image, error) {
	panic("not implemented") // TODO: Implement
}

// Update the image.
func (i *imageService) Update(ctx context.Context, input input.ImageUpdate) error {
	panic("not implemented") // TODO: Implement
}

// Delete an image.
func (i *imageService) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
