package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/thanishsid/goserver/assets"
	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/graphql"
	"github.com/thanishsid/goserver/setup"

	"github.com/thanishsid/goserver/infrastructure/db"
	"github.com/thanishsid/goserver/infrastructure/mailer"
	"github.com/thanishsid/goserver/infrastructure/rediscache"
	"github.com/thanishsid/goserver/infrastructure/tokenizer"
	"github.com/thanishsid/goserver/service"
)

func main() {
	// Read configs.
	config.ReadConfig(".env")

	// Connect to postgresql connection pool and get client.
	pgpool, err := pgxpool.Connect(context.Background(), config.C.PostgresSource)
	if err != nil {
		panic(err)
	}

	// Run postgresql database migrations from sql files in the embedded file system.
	if err := db.Migrate(assets.Files, "sql/migrations"); err != nil {
		panic(err)
	}

	// Initialize database.
	database := db.NewDB(pgpool)

	// Seed the database with default values.
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
		TemplateStore:   assets.Files,
		TemplatePaths:   []string{"mail-templates/*.html"},
	})
	if err != nil {
		panic(err)
	}

	// Create all services.
	imageService := &service.Image{DB: database}

	sessionService := &service.Session{CacheStore: sessiondb}

	userService := &service.User{
		Tokens:         tokenClient,
		Mail:           mailClient,
		DB:             database,
		SessionService: sessionService,
	}

	authService := &service.Auth{
		UserService:    userService,
		SessionService: sessionService,
		GoogleConfig: &oauth2.Config{
			ClientID:     config.C.GoogleOauthClientID,
			ClientSecret: config.C.GoogleOauthClientSecret,
			Endpoint:     google.Endpoint,
			RedirectURL:  config.C.GoogleOauthRedirectURL,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
	}

	// Create new server.
	server := graphql.NewServer(graphql.ServerDeps{
		UserService:    userService,
		ImageService:   imageService,
		SessionService: sessionService,
		AuthService:    authService,
	})

	// Run server.
	log.Fatal(http.ListenAndServe(":"+config.C.ServerPort, server))
}
