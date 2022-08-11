package service

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
)

const TempImageSignature = "temp_image_*"

type Image struct {
	DB db.DB
}

var _ domain.ImageService = (*Image)(nil)

// Reads the image data from a reader and saves it and then returns an image object.
func (i *Image) Save(ctx context.Context, input domain.ImageUploadInput) (*domain.Image, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// Check database to find if image with the same filename and filesize exist.
	fileExists, err := i.DB.CheckImageFileExists(ctx, input.FileName)
	if err != nil {
		return nil, err
	}

	// Return ErrFileAlreadyExists if file exists.
	if fileExists {
		return nil, ErrFileAlreadyExists
	}

	// create a random uuid to use as image filename and the id for database entry.
	imageID := uuid.New()

	if err := input.Validate(); err != nil {
		return nil, err
	}

	// Create a new temporary file to write the file data.
	tempFile, err := os.CreateTemp(config.C.ImageDirectory, TempImageSignature)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	written, err := io.Copy(tempFile, input.File)
	if err != nil {
		return nil, err
	}

	if written != input.Size {
		return nil, ErrIncompleteFile
	}

	now := time.Now()

	newImage := &domain.Image{
		ID:        imageID,
		FileName:  input.FileName,
		Link:      config.C.ImageProxyLink,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save image info to database.
	if err := i.DB.InsertOrUpdateImage(ctx, db.InsertOrUpdateImageParams{
		ID:        newImage.ID,
		FileName:  newImage.FileName,
		CreatedAt: newImage.CreatedAt,
		UpdatedAt: newImage.UpdatedAt,
	}); err != nil {
		return nil, err
	}

	// Rename temp image path to regular image path.
	if err := os.Rename(tempFile.Name(), input.FileName); err != nil {
		return nil, err
	}

	return newImage, nil
}

// Delete an image.
func (i *Image) Delete(ctx context.Context, id uuid.UUID) error {
	image, err := i.DB.GetImageById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	if err := i.DB.DeleteImage(ctx, image.ID); err != nil {
		return err
	}

	return os.Remove(getImagePath(image.FileName))
}

// Get an image.
func (i *Image) One(ctx context.Context, id uuid.UUID) (*domain.Image, error) {
	dbImage, err := i.DB.GetImageById(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &domain.Image{
		ID:        dbImage.ID,
		FileName:  dbImage.FileName,
		Link:      config.C.ImageProxyLink,
		CreatedAt: dbImage.CreatedAt,
		UpdatedAt: dbImage.UpdatedAt,
	}, nil
}

// Get all images in a set of ids.
func (i *Image) AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*domain.Image, error) {
	imageRows, err := i.DB.GetAllImagesInIDS(ctx, ids)
	if err != nil {
		return nil, err
	}

	images := make([]*domain.Image, len(imageRows))

	for i, row := range imageRows {
		images[i] = &domain.Image{
			ID:        row.ID,
			FileName:  row.FileName,
			Link:      config.C.ImageProxyLink,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	}

	return images, nil
}
