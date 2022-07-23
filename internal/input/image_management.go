package input

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type ImageUpload struct {
	Title null.String
	File  io.Reader
}

func (i ImageUpload) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.File, vd.By(func(_ interface{}) error {
			buf := make([]byte, 512)
			nBytes, err := i.File.Read(buf)
			if err != nil {
				return err
			}
			log.Printf("validation read %d bytes from reader\n", nBytes)

			mimeType := http.DetectContentType(buf)

			splitMime := strings.Split(mimeType, "/")

			if splitMime[0] != "image" {
				return errors.New("invalid file type, the uploaded file is not an image")
			}

			return nil
		})),
	)
}

type ImageUpdate struct {
	ID    uuid.UUID
	Title null.String
	File  io.Reader
}

func (i ImageUpdate) Validate() error {
	return vd.ValidateStruct(&i,
		vd.Field(&i.ID, is.UUIDv4),
		vd.Field(&i.File, vd.By(func(_ interface{}) error {
			buf := make([]byte, 512)
			nBytes, err := i.File.Read(buf)
			if err != nil {
				return err
			}
			log.Printf("validation read %d bytes from reader\n", nBytes)

			mimeType := http.DetectContentType(buf)

			splitMime := strings.Split(mimeType, "/")

			if splitMime[0] != "image" {
				return errors.New("invalid file type, the uploaded file is not an image")
			}

			return nil
		})),
	)
}
