package domain

import "gopkg.in/guregu/null.v4"

type ListData[T any] struct {
	Nodes       []*T
	Cursors     []string
	HasMore     bool
	StartCursor null.String
	EndCursor   null.String
}
