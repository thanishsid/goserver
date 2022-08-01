package service

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/thanishsid/goserver/config"
)

var now = time.Now

// Get file path from image directory in config and the image id.
func getImagePath(id uuid.UUID) string {
	return filepath.Join(config.C.ImageDirectory, id.String())
}

// generate a SHA512 hash for the file.
func generateFileHash(bytes []byte) ([]byte, error) {
	hasher := sha512.New()

	_, err := hasher.Write(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate hash of image, %w", err)
	}

	return hasher.Sum(nil), nil
}

// Encode cursor to json and url safe base64.
func EncodeCursor(obj any) (string, error) {
	jsn, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(jsn), nil
}

// Decode cursor from url safe bas64 and json to target type.
func DecodeCursor[T any](cursor string) (T, error) {
	var obj T

	jsn, err := base64.URLEncoding.DecodeString(cursor)
	if err != nil {
		return obj, err
	}

	if err := json.Unmarshal(jsn, &obj); err != nil {
		return obj, err
	}

	return obj, nil
}

// Cleanup, delete image on non nil error.
func CleanupOnError(err *error, path string) {
	if *err != nil {
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			log.Println("failed to cleanup image")
		}
	}
}
