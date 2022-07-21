package search

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/meilisearch/meilisearch-go"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/internal/security"
)

const (
	IndexTypeUsers = "users"
)

type Searcher struct {
	Users UserSearcher
}

func NewSearcher(pg *pgxpool.Pool, searchClient meilisearch.ClientInterface) (*Searcher, error) {

	q := postgres.New(pg)

	userDocs, err := buildUserDocs(q)
	if err != nil {
		return nil, err
	}

	_, err = searchClient.Index(IndexTypeUsers).AddDocuments(userDocs, "id")
	if err != nil {
		return nil, err
	}

	return &Searcher{
		Users: &userSearch{searchClient.Index(IndexTypeUsers)},
	}, nil
}

// Build the list of users from the database for indexing.
func buildUserDocs(q postgres.Querier) ([]domain.User, error) {
	userRows, err := q.GetAllUsers(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	docs := make([]domain.User, len(userRows))

	for i, row := range userRows {
		payload := domain.User{
			ID:        row.ID,
			Email:     row.Email,
			Username:  row.Username,
			FullName:  row.FullName,
			RoleID:    security.Role(row.RoleID),
			PictureID: row.PictureID,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
			DeletedAt: row.DeletedAt,
		}

		docs[i] = payload
	}

	return docs, nil
}
