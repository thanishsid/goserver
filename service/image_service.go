package service

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
)

type Image struct {
	DB db.DB
}

var _ domain.ImageService = (*Image)(nil)

// Reads the image data from a reader and saves it and then returns image object
func (i *Image) Save(ctx context.Context, input domain.ImageUploadInput) (image *domain.Image, err error) {

	// create a random uuid to use as image filename and the id for database entry.
	imageID := uuid.New()

	if err = input.Validate(); err != nil {
		return nil, err
	}

	// Read the image bytes from multipart form.
	imgBytes, err := io.ReadAll(input.File)
	if err != nil {
		return nil, err
	}

	// generate hash from bytes read.
	imageHash, err := generateFileHash(imgBytes)
	if err != nil {
		return nil, err
	}

	// Check if an image with the same hash exists in the database.
	imageExists, err := i.DB.CheckImageHashExists(ctx, imageHash)
	if err != nil {
		return nil, err
	}

	if imageExists {
		return nil, fmt.Errorf("duplicate image, %w", ErrFileAlreadyExists)
	}

	// Create the image directory if it does not exist.
	if err := os.MkdirAll(config.C.ImageDirectory, 0777); err != nil {
		return nil, err
	}

	// Create a new file to store the image.
	dst, err := os.Create(getImagePath(imageID))
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// write image data to the file.
	_, err = dst.Write(imgBytes)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newImage := &domain.Image{
		ID:        imageID,
		Title:     input.Title,
		Link:      config.C.ImageProxyLink,
		FileHash:  imageHash,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save image info to database.
	if err := i.DB.InsertOrUpdateImage(ctx, db.InsertOrUpdateImageParams{
		ID:        newImage.ID,
		Title:     newImage.Title,
		FileHash:  newImage.FileHash,
		CreatedAt: newImage.CreatedAt,
		UpdatedAt: newImage.UpdatedAt,
	}); err != nil {
		return nil, err
	}

	return newImage, nil
}

// Update the image
func (i *Image) Update(ctx context.Context, input domain.ImageUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	// Get existing image info from the database.
	dbImage, err := i.DB.GetImageById(ctx, input.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	// Read the image bytes from multipart form.
	imgBytes, err := io.ReadAll(input.File)
	if err != nil {
		return err
	}

	// Generate hash for the provided image.
	imageHash, err := generateFileHash(imgBytes)
	if err != nil {
		return err
	}

	hashChanged := string(dbImage.FileHash) != string(imageHash)

	if hashChanged {
		// Open the existing image file from disk.
		existingFile, err := os.OpenFile(getImagePath(input.ID), os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer existingFile.Close()

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
		if _, err := existingFile.Write(imgBytes); err != nil {
			return err
		}
	}

	updateParams := db.InsertOrUpdateImageParams(dbImage)

	titleChanged := dbImage.Title != input.Title

	if titleChanged {
		updateParams.Title = input.Title
	}

	if hashChanged {
		updateParams.FileHash = imageHash
	}

	if !titleChanged || !hashChanged {
		return ErrNoChange
	}

	updateParams.UpdatedAt = time.Now()

	return i.DB.InsertOrUpdateImage(ctx, updateParams)
}

// Delete an image
func (i *Image) Delete(ctx context.Context, id uuid.UUID) error {
	if err := os.Remove(getImagePath(id)); err != nil {
		return err
	}

	return i.DB.DeleteImage(ctx, id)
}

// Get an image.
func (i *Image) One(ctx context.Context, id uuid.UUID) (*domain.Image, error) {
	dbImage, err := i.DB.GetImageById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Image{
		ID:        dbImage.ID,
		Title:     dbImage.Title,
		Link:      config.C.ImageProxyLink,
		FileHash:  dbImage.FileHash,
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
			Title:     row.Title,
			Link:      config.C.ImageProxyLink,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	}

	return images, nil
}
