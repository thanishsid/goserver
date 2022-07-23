package search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/internal/input"
)

type UserSearcher interface {
	// Add a user to the search index or update an existing user.
	AddOrUpdateUser(ctx context.Context, user *domain.User) error

	// Remove user with the given id from the search index.
	RemoveUser(ctx context.Context, id uuid.UUID) error

	// Search the index for users based on given params.
	SearchUsers(ctx context.Context, params input.UserFilterBase) ([]domain.User, error)
}

var _ UserSearcher = (*userSearch)(nil)

type userSearch struct {
	meilisearch.IndexInterface
}

// Add a user to the search index or update an existing user.
func (u *userSearch) AddOrUpdateUser(ctx context.Context, user *domain.User) error {

	payload := new(UserPayload)
	payload.LoadFromUser(user)
	docs := []*UserPayload{payload}

	_, err := u.AddDocuments(docs, "id")
	return err
}

// Remove user with the given id from the search index.
func (u *userSearch) RemoveUser(ctx context.Context, id uuid.UUID) error {
	_, err := u.Delete(id.String())
	return err
}

// Search the index for users based on given params.
func (u *userSearch) SearchUsers(ctx context.Context, params input.UserFilterBase) ([]domain.User, error) {

	var filters [][]string

	if params.Role.Valid {
		filters = append(filters, []string{fmt.Sprintf(`role = %s`, params.Role.ValueOrZero())})
	}

	if params.UpdatedAfter.Valid {
		filters = append(filters, []string{fmt.Sprintf(`updatedAt < %d`, params.UpdatedAfter.ValueOrZero().Unix())})
	}

	if !params.ShowDeleted.ValueOrZero() {
		filters = append(filters, []string{"deletedAt = 0"})
	}

	req := &meilisearch.SearchRequest{
		Limit:  params.Limit.ValueOrZero(),
		Filter: filters,
	}

	resp, err := u.Search(params.Query.ValueOrZero(), req)
	if err != nil {
		return nil, err
	}

	usersResult := make([]UserPayload, 0)

	respJsn, err := json.Marshal(resp.Hits)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respJsn, &usersResult); err != nil {
		return nil, err
	}

	users := make([]domain.User, len(usersResult))

	for i, res := range usersResult {
		users[i] = res.ToUser()
	}

	return users, nil
}
