package service

import (
	"testing"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/thanishsid/goserver/infrastructure/db"
	"gopkg.in/guregu/null.v4"
)

func Test_generateFileHash(t *testing.T) {
	t.Parallel()

	randbytes := []byte("abcdefghijk")

	hash, err := generateFileHash(randbytes)
	require.NoError(t, err)
	require.Len(t, hash, 512/8)
}

func Test_encode_decode_Cursor(t *testing.T) {
	t.Parallel()

	obj := db.GetManyUsersParams{
		Search:        null.StringFrom(fake.Name()),
		Role:          null.StringFrom(getRandRole().String()),
		UpdatedBefore: null.TimeFrom(fake.Date()),
		ShowDeleted:   true,
		UsersLimit:    50,
	}

	cursorStr, err := encodeCursor(obj)
	require.NoError(t, err)
	require.NotEmpty(t, cursorStr)

	decodedObj, err := decodeCursor[db.GetManyUsersParams](cursorStr)
	require.NoError(t, err)
	require.Equal(t, decodedObj, obj)
}
