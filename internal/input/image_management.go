package input

import (
	"io"

	"github.com/google/uuid"
)

type ImageUpload struct {
	Title string    `json:"title"`
	File  io.Reader `json:"file"`
}

type ImageUpdate struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	File  io.Reader `json:"file"`
}
