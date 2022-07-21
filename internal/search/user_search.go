package search

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
	"gopkg.in/guregu/null.v4"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/internal/security"
)

type UserSearcher interface {
	// Add a user to the search index or update an existing user.
	AddOrUpdateUser(ctx context.Context, user *domain.User) error

	// Remove user with the given id from the search index.
	RemoveUser(ctx context.Context, id uuid.UUID) error

	// Search the index for users based on given params.
	SearchUsers(ctx context.Context, params UserSearchParams) ([]domain.User, error)
}

var _ UserSearcher = (*userSearch)(nil)

type UserSearchParams struct {
	SearchPhrase null.String
	Role         security.Role
	ShowDeleted  bool
	Limit        int64
	Offset       int64
}

type userSearch struct {
	meilisearch.IndexInterface
}

// Add a user to the search index or update an existing user.
func (u *userSearch) AddOrUpdateUser(ctx context.Context, user *domain.User) error {
	docs := []*domain.User{user}

	_, err := u.AddDocuments(docs, "id")
	return err
}

// Remove user with the given id from the search index.
func (u *userSearch) RemoveUser(ctx context.Context, id uuid.UUID) error {
	_, err := u.Delete(id.String())
	return err
}

// Search the index for users based on given params.
func (u *userSearch) SearchUsers(ctx context.Context, params UserSearchParams) ([]domain.User, error) {
	resp, err := u.Search(params.SearchPhrase.ValueOrZero(), &meilisearch.SearchRequest{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}

	users := make([]domain.User, 0)

	respJsn, err := json.Marshal(resp.Hits)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respJsn, &users); err != nil {
		return nil, err
	}

	return users, nil
}
