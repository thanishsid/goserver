package ftsearch

import (
	"errors"

	"github.com/meilisearch/meilisearch-go"

	"github.com/thanishsid/goserver/config"
)

func NewSearchClient() (meilisearch.ClientInterface, error) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: config.C.MeilisearchSource,
	})

	if client.IsHealthy() {
		return client, nil
	}

	return nil, errors.New("meilisearch connection failed")
}
