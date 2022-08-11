package domain

import (
	"io"

	"gopkg.in/guregu/null.v4"
)

type ContextKey string

type ListData[T any] struct {
	Nodes       []*T
	Cursors     []string
	HasMore     bool
	StartCursor null.String
	EndCursor   null.String
}

type FileUploadData struct {
	File        io.Reader
	FileName    string
	Size        int64
	ContentType string
}
