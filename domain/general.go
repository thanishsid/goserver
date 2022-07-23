package domain

import "gopkg.in/guregu/null.v4"

type ListWithCursor[T any] struct {
	Data       []T         `json:"data"`
	NextCursor null.String `json:"nextCursor"`
}
