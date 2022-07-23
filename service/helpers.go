package service

import (
	"crypto/sha512"
	"fmt"
	"io"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/thanishsid/goserver/config"
)

// Get file path from image directory in config and the image id.
func getImagePath(id uuid.UUID) string {
	return filepath.Join(config.C.ImageDirectory, id.String())
}

// generate a SHA512 hash for the file.
func generateFileHash(file io.Reader) (string, error) {
	hasher := sha512.New()

	_, err := io.Copy(hasher, file)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash of image, %w", err)
	}

	return string(hasher.Sum(nil)), nil
}
