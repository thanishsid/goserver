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

type VideoUpload struct {
	Title             null.String
	File              io.Reader
	CustomThumbnailID uuid.NullUUID
}

func (v VideoUpload) Validate() error {
	return vd.ValidateStruct(&v,
		vd.Field(&v.File, vd.By(func(_ interface{}) error {
			buf := make([]byte, 512)
			nBytes, err := v.File.Read(buf)
			if err != nil {
				return err
			}
			log.Printf("validation read %d bytes from reader\n", nBytes)

			mimeType := http.DetectContentType(buf)

			splitMime := strings.Split(mimeType, "/")

			if splitMime[0] != "video" {
				return errors.New("invalid file type, the uploaded file is not an video")
			}

			return nil
		})),
	)
}

type VideoUpdate struct {
	ID                uuid.UUID
	Title             null.String
	File              io.Reader
	CustomThumbnailID uuid.NullUUID
}

func (v VideoUpdate) Validate() error {
	return vd.ValidateStruct(&v,
		vd.Field(&v.File, is.UUIDv4),
		vd.Field(&v.File, vd.By(func(_ interface{}) error {
			buf := make([]byte, 512)
			nBytes, err := v.File.Read(buf)
			if err != nil {
				return err
			}
			log.Printf("validation read %d bytes from reader\n", nBytes)

			mimeType := http.DetectContentType(buf)

			splitMime := strings.Split(mimeType, "/")

			if splitMime[0] != "video" {
				return errors.New("invalid file type, the uploaded file is not an video")
			}

			return nil
		})),
	)
}
