package config

import (
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joho/godotenv"
	"github.com/thanishsid/goserver/domain"
)

type config struct {
	CompanyName string

	PostgresSource     string
	RedisSessionSource string

	AdminEmail string

	ImageDirectory string
	ImageProxyLink string

	MailHost     string
	MailPort     int
	MailUser     string
	MailPassword string

	OauthState string

	GoogleOauthClientID     string
	GoogleOauthClientSecret string
	GoogleOauthRedirectURL  string

	TokenSigningKey    string
	TokenEncryptionKey string

	ServerPort string

	Environment string
}

func (c config) Validate() error {
	return vd.ValidateStruct(&c,
		vd.Field(&c.CompanyName, vd.Required),

		vd.Field(&c.PostgresSource, vd.Required),
		vd.Field(&c.RedisSessionSource, vd.Required),

		vd.Field(&c.AdminEmail, vd.Required, domain.IsEmail),

		vd.Field(&c.ImageDirectory, vd.Required),
		vd.Field(&c.ImageProxyLink, vd.Required),

		vd.Field(&c.MailHost, vd.Required),
		vd.Field(&c.MailPort, vd.Required),
		vd.Field(&c.MailUser, vd.Required),
		vd.Field(&c.MailPassword, vd.Required),

		vd.Field(&c.OauthState, vd.Required),
		vd.Field(&c.GoogleOauthClientID, vd.Required),
		vd.Field(&c.GoogleOauthClientSecret, vd.Required),
		vd.Field(&c.GoogleOauthRedirectURL, vd.Required),

		vd.Field(&c.TokenSigningKey, vd.Required),
		vd.Field(&c.TokenEncryptionKey, vd.Required),

		vd.Field(&c.ServerPort, vd.Required),
		vd.Field(&c.Environment, vd.Required),
	)
}

var C config

func ReadConfig(files ...string) {
	if err := godotenv.Load(files...); err != nil {
		panic(err)
	}

	C = config{
		CompanyName: os.Getenv("COMPANY_NAME"),

		PostgresSource:     os.Getenv("POSTGRES_SOURCE"),
		RedisSessionSource: os.Getenv("SESSION_DB_SOURCE"),

		AdminEmail: os.Getenv("ADMIN_EMAIL"),

		ImageDirectory: os.Getenv("IMAGE_DIRECTORY"),
		ImageProxyLink: os.Getenv("IMGPROXY_LINK"),

		MailHost:     os.Getenv("MAIL_HOST"),
		MailPort:     MustParseInt[int](os.Getenv("MAIL_PORT")),
		MailUser:     os.Getenv("MAIL_USER"),
		MailPassword: os.Getenv("MAIL_PASSWORD"),

		OauthState:              os.Getenv("OAUTH_STATE"),
		GoogleOauthClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		GoogleOauthClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		GoogleOauthRedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),

		TokenSigningKey:    os.Getenv("TOKEN_SIGNING_KEY"),
		TokenEncryptionKey: os.Getenv("TOKEN_ENCRYPTION_KEY"),

		ServerPort:  os.Getenv("SERVER_PORT"),
		Environment: os.Getenv("ENVIRONMENT"),
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
