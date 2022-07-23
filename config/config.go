package config

import (
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joho/godotenv"
)

type config struct {
	CompanyName        string
	PostgresSource     string
	RedisSessionSource string
	MeilisearchSource  string
	ImageDirectory     string
	ImageProxyLink     string

	MailHost     string
	MailPort     int
	MailUser     string
	MailPassword string

	GoogleOauthClientID     string
	GoogleOauthClientSecret string

	TokenSigningKey    string
	TokenEncryptionKey string

	Environment string
}

func (c config) Validate() error {
	return vd.ValidateStruct(&c,
		vd.Field(&c.CompanyName, vd.Required),
		vd.Field(&c.PostgresSource, vd.Required),
		vd.Field(&c.RedisSessionSource, vd.Required),
		vd.Field(&c.MeilisearchSource, vd.Required),
		vd.Field(&c.ImageDirectory, vd.Required),
		vd.Field(&c.ImageProxyLink, vd.Required),
		vd.Field(&c.MailHost, vd.Required),
		vd.Field(&c.MailPort, vd.Required),
		vd.Field(&c.MailUser, vd.Required),
		vd.Field(&c.MailPassword, vd.Required),
		vd.Field(&c.GoogleOauthClientID, vd.Required),
		vd.Field(&c.GoogleOauthClientSecret, vd.Required),
		vd.Field(&c.TokenSigningKey, vd.Required),
		vd.Field(&c.TokenEncryptionKey, vd.Required),
	)
}

var C config

func ReadConfig(files ...string) {
	if err := godotenv.Load(files...); err != nil {
		panic(err)
	}

	C = config{
		CompanyName:             "Golang Corp",
		PostgresSource:          os.Getenv("POSTGRES_SOURCE"),
		RedisSessionSource:      os.Getenv("SESSION_DB_SOURCE"),
		MeilisearchSource:       os.Getenv("MEILISEARCH_SOURCE"),
		ImageDirectory:          os.Getenv("IMAGE_DIRECTORY"),
		ImageProxyLink:          os.Getenv("IMGPROXY_LINK"),
		MailHost:                os.Getenv("MAIL_HOST"),
		MailPort:                MustParseInt[int](os.Getenv("MAIL_PORT")),
		MailUser:                os.Getenv("MAIL_USER"),
		MailPassword:            os.Getenv("MAIL_PASSWORD"),
		GoogleOauthClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		GoogleOauthClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		TokenSigningKey:         os.Getenv("TOKEN_SIGNING_KEY"),
		TokenEncryptionKey:      os.Getenv("TOKEN_ENCRYPTION_KEY"),
		Environment:             os.Getenv("ENVIRONMENT"),
	}

	if err := C.Validate(); err != nil {
		panic(err)
	}

	spew.Dump(C)
}

func MustParseInt[T int | int64 | int32](numStr string) T {
	parsedInt, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return T(parsedInt)
}
