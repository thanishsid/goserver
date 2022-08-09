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

	"github.com/thanishsid/goserver/infrastructure/db"
	"github.com/thanishsid/goserver/infrastructure/mailer"
	"github.com/thanishsid/goserver/infrastructure/rediscache"
	"github.com/thanishsid/goserver/infrastructure/security"
	"github.com/thanishsid/goserver/infrastructure/tokenizer"
	"github.com/thanishsid/goserver/service"
)

func main() {
	// Read configs.
	config.ReadConfig(".env")

	// Connect to postgresql connection pool and get client.
	pgpool, err := pgxpool.Connect(context.Background(), config.C.PostgresSource)
	PanicOnError(err)
	defer pgpool.Close()
	PanicOnError(pgpool.Ping(context.Background()))

	// Run postgresql database migrations from sql files in the embedded file system.
	PanicOnError(db.Migrate(assets.Files, "sql/migrations"))

	// Initialize database.
	database := db.NewDB(pgpool)

	// Seed the database with default values.
	PanicOnError(db.Seed(context.Background(), database))

	// Connect to redis session server with LFU in process cache.
	sessiondb, err := rediscache.NewCacheStore(config.C.RedisSessionSource, 0)
	PanicOnError(err)

	// Create new token client.
	tokenClient, err := tokenizer.NewTokenizer()
	PanicOnError(err)

	// Connect and obtain mail client.
	mailClient, err := mailer.NewMailer(mailer.MailerConfig{
		DialerTimeout:   time.Second * 15,
		DialerTLSConfig: &tls.Config{InsecureSkipVerify: config.C.Environment != "production"},
		TemplateStore:   assets.Files,
		TemplatePaths:   []string{"mail-templates/*.html"},
	})
	PanicOnError(err)

	// Create all services
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

	authorizer, err := security.NewAuthorizer(assets.Files)
	PanicOnError(err)

	// Create new server
	server := graphql.NewServer(graphql.ServerDeps{
		UserService:    userService,
		ImageService:   imageService,
		SessionService: sessionService,
		AuthService:    authService,
		Authorizer:     authorizer,
	})

	// Run server.
	log.Fatal(http.ListenAndServe(":"+config.C.ServerPort, server))
}

// Check if passed error is nil if not will panic.
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
