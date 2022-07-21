package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://localhost:7700",
	})

	if !client.IsHealthy() {
		log.Fatal("search client unhealthy")
	}

	client.Index("users").DeleteAllDocuments()

	users := []User{
		{
			ID:        uuid.New(),
			Name:      "Thanish",
			Email:     "thanish@gmail.com",
			CreatedAt: time.Now().Add(-time.Hour * 24 * 300),
		},
		{
			ID:        uuid.New(),
			Name:      "Dominic",
			Email:     "dominic@gmail.com",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Name:      "Gaveen",
			Email:     "gaveen@gmail.com",
			CreatedAt: time.Now().Add(-time.Hour * 24 * 3),
		},
	}

	resp, err := client.Index("users").AddDocuments(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n\n", *resp)

	time.Sleep(time.Second * 1)

	settings := meilisearch.Settings{
		SortableAttributes: []string{
			"createdAt",
		},
	}

	_, err = client.Index("users").UpdateSettings(&settings)
	if err != nil {
		log.Fatal(err)
	}

	sortables, err := client.Index("users").GetSortableAttributes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", *sortables)

	time.Sleep(time.Second * 1)

	searchResp, err := client.Index("users").Search("", &meilisearch.SearchRequest{
		Sort: []string{
			"createdAt:desc",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	var results []User

	respJsn, _ := json.Marshal(searchResp.Hits)

	if err := json.Unmarshal(respJsn, &results); err != nil {
		log.Fatal(err)
	}

	for _, res := range results {
		fmt.Printf("%+v\n", res)
	}

}
