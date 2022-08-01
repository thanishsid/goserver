package resolver

import (
	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

func UUIDFromPtr(strPtr *string) uuid.NullUUID {
	if strPtr == nil {
		return uuid.NullUUID{}
	}

	id, err := uuid.Parse(*strPtr)

	return uuid.NullUUID{UUID: id, Valid: err == nil}
}

type Ints interface {
	int | int8 | int16 | int32 | int64
}

func NullIntFromPtr[T Ints](intPtr *T) null.Int {
	if intPtr != nil {
		return null.IntFrom(int64(*intPtr))
	}

	return null.Int{}
}
