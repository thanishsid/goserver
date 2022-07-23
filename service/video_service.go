package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/input"
)

type videoService struct {
	*ServiceDeps
}

var _ domain.VideoService = (*videoService)(nil)

// Reads the video data from reader, saves it to disk and returns video object.
func (v *videoService) Save(ctx context.Context, input input.VideoUpload) (*domain.Video, error) {
	panic("not implemented") // TODO: Implement
}

// Update the video.
func (v *videoService) Update(ctx context.Context, input input.VideoUpdate) error {
	panic("not implemented") // TODO: Implement
}

// Delete a video.
func (v *videoService) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

// Get Videos.
func (v *videoService) Videos(ctx context.Context, filter input.MediaFilter) ([]domain.Video, error) {
	panic("not implemented") // TODO: Implement
}
