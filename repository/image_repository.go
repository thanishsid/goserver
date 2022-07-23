package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/internal/input"
)

type imageRepository struct {
	db postgres.Querier
}

var _ domain.ImageRepository = (*imageRepository)(nil)

func (i *imageRepository) SaveOrUpdate(ctx context.Context, image *domain.Image) error {

	image.Link = config.C.ImageProxyLink

	if err := image.Validate(); err != nil {
		return err
	}

	return i.db.InsertOrUpdateImage(ctx, postgres.InsertOrUpdateImageParams{
		ID:        image.ID,
		Title:     image.Title,
		FileHash:  image.FileHash,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	})
}

func (i *imageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return i.db.DeleteImage(ctx, id)
}

func (i *imageRepository) OneByID(ctx context.Context, id uuid.UUID) (*domain.Image, error) {
	imageRow, err := i.db.GetImageById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Image{
		ID:        imageRow.ID,
		Title:     imageRow.Title,
		Link:      config.C.ImageProxyLink,
		FileHash:  imageRow.FileHash,
		CreatedAt: imageRow.CreatedAt,
		UpdatedAt: imageRow.UpdatedAt,
	}, nil
}

func (i *imageRepository) All(ctx context.Context, ids ...uuid.UUID) ([]domain.Image, error) {
	imageRows, err := i.db.GetAllImages(ctx, ids)
	if err != nil {
		return nil, err
	}

	images := make([]domain.Image, len(imageRows))

	for i, row := range imageRows {
		images[i] = domain.Image{
			ID:        row.ID,
			Title:     row.Title,
			Link:      config.C.ImageProxyLink,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.CreatedAt,
		}
	}

	return images, nil
}

func (i *imageRepository) Many(ctx context.Context, filter input.MediaFilterBase) ([]domain.Image, error) {
	imageRows, err := i.db.GetManyImages(ctx, postgres.GetManyImagesParams{
		ViewUnused:   filter.ViewUnused,
		UpdatedAfter: filter.UpdatedAfter,
		ImageLimit:   filter.Limit.ValueOrZero(),
	})
	if err != nil {
		return nil, err
	}

	images := make([]domain.Image, len(imageRows))

	for i, row := range imageRows {
		images[i] = domain.Image{
			ID:        row.ID,
			Title:     row.Title,
			Link:      config.C.ImageProxyLink,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	}

	return images, nil
}

func (i *imageRepository) CheckHashExists(ctx context.Context, hash string) (bool, error) {
	return i.db.CheckImageHashExists(ctx, hash)
}
