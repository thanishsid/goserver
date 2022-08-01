package main

import (
	"context"
	"crypto/tls"
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/graphql"
	"github.com/thanishsid/goserver/setup"

	"github.com/thanishsid/goserver/infrastructure/db"
	"github.com/thanishsid/goserver/infrastructure/mailer"
	"github.com/thanishsid/goserver/infrastructure/rediscache"
	"github.com/thanishsid/goserver/infrastructure/tokenizer"
	"github.com/thanishsid/goserver/service"
)

//go:embed assets sql/migrations
var assets embed.FS

func main() {
	// Read configs.
	config.ReadConfig(".env")

	// Connect to postgresql connection pool and get client.
	pgpool, err := pgxpool.Connect(context.Background(), config.C.PostgresSource)
	if err != nil {
		panic(err)
	}

	// Run postgresql database migrations from sql files in the embedded file system.
	if err := db.Migrate(assets, "sql/migrations"); err != nil {
		panic(err)
	}

	database := db.NewDB(pgpool)

	seeder := setup.NewSeeder(database)

	if err := seeder.UpdateRoles(context.Background()); err != nil {
		panic(err)
	}

	// Connect to redis session server with LFU in process cache.
	sessiondb, err := rediscache.NewCacheStore(config.C.RedisSessionSource, 0)
	if err != nil {
		panic(err)
	}

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

	// Create all services.
	userService := service.NewUserService(tokenClient, mailClient, database)
	imageService := service.NewImageService(database)
	sessionService := service.NewSessionService(sessiondb)

	// Create new server.
	srv := graphql.NewServer(graphql.ServerDeps{
		UserService:    userService,
		ImageService:   imageService,
		SessionService: sessionService,
	})

	// Run server.
	log.Fatal(http.ListenAndServe(":"+config.C.ServerPort, srv))
}
