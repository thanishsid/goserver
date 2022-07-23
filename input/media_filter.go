package input

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"gopkg.in/guregu/null.v4"
)

type MediaFilter struct {
	MediaFilterBase
	Cursor null.String `schema:"cursor"`
}

func (i MediaFilter) GetFilterFromCursor() (MediaFilterBase, error) {
	var imgFilter MediaFilterBase

	if !i.Cursor.Valid {
		return imgFilter, errors.New("unable to get filter from null cursor")
	}

	jsn := make([]byte, len(i.Cursor.String))

	_, err := base64.URLEncoding.Decode(jsn, []byte(i.Cursor.String))
	if err != nil {
		return imgFilter, err
	}

	if err := json.Unmarshal(jsn, &imgFilter); err != nil {
		return imgFilter, err
	}

	return imgFilter, nil
}

type MediaFilterBase struct {
	ViewUnused   bool      `schema:"viewUnused" json:"viewUnused"`
	UpdatedAfter null.Time `schema:"updatedAfter" json:"updatedAfter"`
	Limit        null.Int  `schema:"limit" json:"limit"`
}

func (i MediaFilterBase) CreateCursor() (string, error) {
	jsn, err := json.Marshal(i)
	if err != nil {
		return "", err
	}

	encString := base64.URLEncoding.EncodeToString(jsn)

	return encString, nil
}
