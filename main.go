package main

import (
	"context"
	"crypto/tls"
	"embed"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/infrastructure/ftsearch"
	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/infrastructure/rediscache"
	"github.com/thanishsid/goserver/internal/mailer"
	"github.com/thanishsid/goserver/internal/search"
	"github.com/thanishsid/goserver/internal/sessions"
	"github.com/thanishsid/goserver/internal/tokenizer"
	"github.com/thanishsid/goserver/repository"
	"github.com/thanishsid/goserver/service"
)

//go:embed assets sql/migrations
var assets embed.FS

func main() {
	// Read configs.
	config.ReadConfig(".env")

	// Connect to postgresql connection pool and get client.
	dbpool, err := pgxpool.Connect(context.Background(), config.C.PostgresSource)
	if err != nil {
		panic(err)
	}

	// Run postgresql database migrations from sql files in the embedded file system.
	if err := postgres.Migrate(assets, "sql/migrations"); err != nil {
		panic(err)
	}

	// Connect to redis session server with LFU in process cache.
	sessiondb, err := rediscache.NewCacheStore(config.C.RedisSessionSource, 0)
	if err != nil {
		panic(err)
	}

	// Create new session manager using redis cache.
	sessionMgr := sessions.NewSessionManager(sessiondb, "sessions")

	// Create new token client.
	tokenClient, err := tokenizer.NewTokenizer()
	if err != nil {
		panic(err)
	}

	// Connect and obtain mail client.
	mailClient, err := mailer.NewMailer(mailer.MailerConfig{
		DialerTimeout:   time.Second * 15,
		DialerTLSConfig: &tls.Config{InsecureSkipVerify: config.C.Environment != "production"},
		TemplateStore:   assets,
		TemplatePaths:   []string{"assets/mail-templates/*.html"},
	})
	if err != nil {
		panic(err)
	}

	// Connect to meilisearch and get client.
	searchClient, err := ftsearch.NewSearchClient()
	if err != nil {
		panic(err)
	}

	// Create new search indexer.
	searcher, err := search.NewSearcher(dbpool, searchClient)
	if err != nil {
		panic(err)
	}

	// Create new repository using postgres db connection pool and search indexer.
	rp := repository.New(dbpool, searcher)

	// Create all services.
	service := service.New(&service.ServiceDeps{
		Tokens: tokenClient,
		Mail:   mailClient,
		Repo:   rp,
	})

	fmt.Println(sessionMgr)
	fmt.Println(service)
}
