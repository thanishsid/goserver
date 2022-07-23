package search

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/meilisearch/meilisearch-go"

	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/infrastructure/security"
)

const (
	IndexTypeUsers = "users"
)

type Searcher struct {
	Users UserSearcher
}

func NewSearcher(pg *pgxpool.Pool, searchClient *meilisearch.Client) (*Searcher, error) {

	if !searchClient.IsHealthy() {
		return nil, errors.New("meilisearch connection failed")
	}

	q := postgres.New(pg)

	if err := buildUserIndex(q, searchClient); err != nil {
		return nil, err
	}

	return &Searcher{
		Users: &userSearch{searchClient.Index(IndexTypeUsers)},
	}, nil
}

// Build the list of users from the database for indexing.
func buildUserIndex(q postgres.Querier, sc meilisearch.ClientInterface) error {

	_, err := sc.CreateIndex(&meilisearch.IndexConfig{
		Uid:        IndexTypeUsers,
		PrimaryKey: "id",
	})
	if err != nil {
		return err
	}

	_, err = sc.Index(IndexTypeUsers).UpdateSearchableAttributes(&[]string{
		"email",
		"username",
		"fullName",
		"role",
	})
	if err != nil {
		return err
	}

	_, err = sc.Index(IndexTypeUsers).UpdateRankingRules(&[]string{
		"words",
		"typo",
		"proximity",
		"attribute",
		"sort",
		"exactness",
		"updatedAt:desc",
	})
	if err != nil {
		return err
	}

	_, err = sc.Index(IndexTypeUsers).UpdateFilterableAttributes(&[]string{
		"role",
		"updatedAt",
		"deleteAt",
	})

	userRows, err := q.GetAllUsers(context.Background(), nil)
	if err != nil {
		return err
	}

	docs := make([]UserPayload, len(userRows))

	for i, row := range userRows {
		payload := UserPayload{
			ID:        row.ID,
			Email:     row.Email,
			Username:  row.Username,
			FullName:  row.FullName,
			Role:      security.Role(row.UserRole),
			PictureID: row.PictureID,
			CreatedAt: row.CreatedAt.Unix(),
			UpdatedAt: row.UpdatedAt.Unix(),
		}

		if row.DeletedAt.Valid {
			payload.DeletedAt = row.DeletedAt.ValueOrZero().Unix()
		}

		docs[i] = payload
	}

	_, err = sc.Index(IndexTypeUsers).AddDocuments(docs, "id")
	if err != nil {
		return err
	}

	return nil
}
