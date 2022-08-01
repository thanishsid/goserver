package dataloader

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"

	"github.com/thanishsid/goserver/domain"
)

type userBatcher struct {
	domain.UserService
}

func (b *userBatcher) Get(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	fmt.Printf("dataloader.userBatcher.get, users: [%s]\n", strings.Join(keys.Keys(), ","))

	// create a map for remembering the order of keys passed in
	keyOrder := make(map[uuid.UUID]int, len(keys))

	var userIDs []uuid.UUID

	for ix, key := range keys {

		userID, err := uuid.Parse(key.String())
		if err != nil {
			return []*dataloader.Result{{Data: nil, Error: err}}
		}

		userIDs = append(userIDs, userID)
		keyOrder[userID] = ix
	}

	users, err := b.AllByIDS(ctx, userIDs...)
	if err != nil {
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	results := make([]*dataloader.Result, len(keys))

	for _, user := range users {
		ix, ok := keyOrder[user.ID]

		if ok {
			results[ix] = &dataloader.Result{Data: user, Error: nil}
			delete(keyOrder, user.ID)
		}
	}

	for userID, ix := range keyOrder {
		err := fmt.Errorf("user not found %s", userID)
		results[ix] = &dataloader.Result{Data: nil, Error: err}
	}

	return results
}
