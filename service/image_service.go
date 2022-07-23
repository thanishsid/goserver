package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/internal/input"
)

type imageService struct {
	*ServiceDeps
}

var _ domain.ImageService = (*imageService)(nil)

// Reads the image data from a reader and saves it and then returns image object
func (i *imageService) Save(ctx context.Context, input input.ImageUpload) (*domain.Image, error) {

	if err := input.Validate(); err != nil {
		return nil, err
	}

	// Generate a hash for the image.
	imageHash, err := generateFileHash(input.File)
	if err != nil {
		return nil, err
	}

	// Check if an image with the same hash exists in the database.
	imageExists, err := i.Repo.ImageRepository().CheckHashExists(ctx, imageHash)
	if err != nil {
		return nil, err
	}

	if imageExists {
		return nil, fmt.Errorf("duplicate image, %w", ErrFileAlreadyExists)
	}

	imageID := uuid.New()

	// Create a new file to store the image.
	dst, err := os.Create(getImagePath(imageID))
	defer dst.Close()
	if err != nil {
		return nil, err
	}

	// write image data to the file.
	_, err = io.Copy(dst, input.File)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	newImage := &domain.Image{
		ID:        imageID,
		Title:     input.Title,
		FileHash:  imageHash,
		CreatedAt: now,
		UpdatedAt: now,
	}
	newImage.LoadImageProxyLink()

	// Save new image info to the database.
	if err := i.Repo.ImageRepository().SaveOrUpdate(ctx, newImage); err != nil {
		return nil, err
	}

	return newImage, nil
}

// Update the image
func (i *imageService) Update(ctx context.Context, input input.ImageUpdate) error {
	if err := input.Validate(); err != nil {
		return err
	}

	r := i.Repo.ImageRepository()

	// Get existing image info from the database.
	existingImage, err := r.OneByID(ctx, input.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	// Generate hash the provided image.
	imageHash, err := generateFileHash(input.File)
	if err != nil {
		return err
	}

	updatedImage := *existingImage
	updatedImage.FileHash = imageHash
	updatedImage.Title = input.Title

	if updatedImage.IsEqual(existingImage) {
		return ErrNoChange
	}

	// Open the existing image file from disk.
	existingFile, err := os.OpenFile(getImagePath(input.ID), os.O_RDWR, 0666)
	defer existingFile.Close()
	if err != nil {
		return err
	}

	// Truncate the image file size to 0.
	if err := existingFile.Truncate(0); err != nil {
		return err
	}

	// Reset the io offset to the start of the file.
	_, err = existingFile.Seek(0, 0)
	if err != nil {
		return err
	}

	// Write new image file data to the existing file.
	if _, err := io.Copy(existingFile, input.File); err != nil {
		return err
	}

	// Save the updated image info to the database.
	updatedImage.UpdatedAt = time.Now()
	if err := r.SaveOrUpdate(ctx, &updatedImage); err != nil {
		return err
	}

	return nil
}

// Delete an image
func (i *imageService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := os.Remove(getImagePath(id)); err != nil {
		return err
	}

	return i.Repo.ImageRepository().Delete(ctx, id)
}

// Get images.
func (i *imageService) Images(ctx context.Context, filter input.MediaFilter) (*domain.ListWithCursor[domain.Image], error) {

	var baseFilter input.MediaFilterBase
	var err error

	if filter.Cursor.Valid {
		baseFilter, err = filter.GetFilterFromCursor()
		if err != nil {
			return nil, err
		}
	} else {
		baseFilter = filter.MediaFilterBase
	}

	if baseFilter.Limit.ValueOrZero() == 0 {
		baseFilter.Limit = null.IntFrom(config.DEFAULT_IMAGES_LIST_LIMIT)
	}

	r := i.Repo.ImageRepository()

	// Limit incremented by 1 to find if next page exists based on
	// whether the returned array size is equal to the speculation limit.
	speculationLimit := baseFilter.Limit.ValueOrZero() + 1

	images, err := r.Many(ctx, input.MediaFilterBase{
		ViewUnused:   baseFilter.ViewUnused,
		UpdatedAfter: baseFilter.UpdatedAfter,
		Limit:        null.IntFrom(speculationLimit),
	})
	if err != nil {
		return nil, err
	}

	if len(images) < int(speculationLimit) {
		return &domain.ListWithCursor[domain.Image]{
			Data: images,
		}, nil
	}

	nextCursor, err := baseFilter.CreateCursor()
	if err != nil {
		return nil, err
	}

	return &domain.ListWithCursor[domain.Image]{
		Data:       images[:baseFilter.Limit.ValueOrZero()],
		NextCursor: null.StringFrom(nextCursor),
	}, nil
}
