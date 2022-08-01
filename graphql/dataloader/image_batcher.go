package dataloader

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"

	"github.com/thanishsid/goserver/domain"
)

type imageBatcher struct {
	domain.ImageService
}

func (b *imageBatcher) Get(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	fmt.Printf("dataloader.imageBatcher.get, images: [%s]\n", strings.Join(keys.Keys(), ","))

	keyOrder := make(map[uuid.UUID]int, len(keys))

	var imageIDs []uuid.UUID

	for ix, key := range keys {

		imageID, err := uuid.Parse(key.String())
		if err != nil {
			return []*dataloader.Result{{Data: nil, Error: err}}
		}

		imageIDs = append(imageIDs, imageID)
		keyOrder[imageID] = ix
	}

	images, err := b.AllByIDS(ctx, imageIDs...)
	if err != nil {
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	results := make([]*dataloader.Result, len(keys))

	for _, image := range images {
		ix, ok := keyOrder[image.ID]

		if ok {
			results[ix] = &dataloader.Result{Data: image, Error: nil}
			delete(keyOrder, image.ID)
		}
	}

	for imageID, ix := range keyOrder {
		err := fmt.Errorf("image not found %s", imageID)
		results[ix] = &dataloader.Result{Data: nil, Error: err}
	}

	return results
}
