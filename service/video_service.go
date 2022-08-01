package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/thanishsid/goserver/domain"
)

type videoService struct{}

var _ domain.VideoService = (*videoService)(nil)

// Reads the video data from reader, saves it to disk and returns video object.
func (v *videoService) Save(ctx context.Context, input domain.VideoUploadInput) (*domain.Video, error) {
	panic("not implemented") // TODO: Implement
}

// Update the video.
func (v *videoService) Update(ctx context.Context, input domain.VideoUpdateInput) error {
	panic("not implemented") // TODO: Implement
}

// Delete a video.
func (v *videoService) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

// Get Videos by ids.
func (v *videoService) AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*domain.Video, error) {
	panic("not implemented") // TODO: Implement
}
