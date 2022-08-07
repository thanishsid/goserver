package resolver

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/thanishsid/goserver/config"
	null "gopkg.in/guregu/null.v4"
)

// Get NullUUID from string ptr.
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

// Get null.Int from int[bitsize] ptr.
func NullIntFromPtr[T Ints](intPtr *T) null.Int {
	if intPtr != nil {
		return null.IntFrom(int64(*intPtr))
	}

	return null.Int{}
}

// Get UserAgent from context.
func UseragentFor(ctx context.Context) (string, error) {
	ua, ok := ctx.Value(config.USERAGENT_KEY).(string)
	if !ok {
		return "", errors.New("useragnt not found in context")
	}

	return ua, nil
}
