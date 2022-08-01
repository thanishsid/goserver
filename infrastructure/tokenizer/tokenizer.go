package tokenizer

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/thanishsid/goserver/config"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var ErrNonPointerClaim = errors.New("invalid claims type, the claims must be a pointer")

type ValidationError = vd.Errors

type Validateable interface {
	Validate() error
}

type TokenConfig struct {
	SigningKey    string
	EncryptionKey string
}

func NewTokenizer() (Tokenizer, error) {

	signingKey := []byte(config.C.TokenSigningKey)
	encryptionKey := []byte(config.C.TokenEncryptionKey)

	if len(signingKey) == 0 {
		return nil, errors.New("empty token signing key")
	}

	if len(encryptionKey) == 0 {
		return nil, errors.New("empty token encryption key")
	}

	return &tokenizer{
		SigningKey:    signingKey,
		EncryptionKey: encryptionKey,
	}, nil
}

type Tokenizer interface {
	// Create a new encrypted token.
	CreateToken(ctx context.Context, claims Validateable) (string, error)

	// Takes the encrypted token and a pointer to the claims, the token is verified and deserialized into the provided claims pointer.
	GetClaims(ctx context.Context, secureToken string, claims Validateable) error

	// Takes the encrypted token and a pointer to the claims, the token is deserialized into the provided claims pointer.
	GetClaimsUnsafe(ctx context.Context, secureToken string, claims Validateable) error
}

type tokenizer struct {
	SigningKey    []byte
	EncryptionKey []byte
}

// Create a new encrypted token.
func (tc *tokenizer) CreateToken(ctx context.Context, claims Validateable) (string, error) {

	var token string

	if err := claims.Validate(); err != nil {
		return token, fmt.Errorf("failed to create new token, invalid claims: %w", err)
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{
			Algorithm: jose.HS256,
			Key:       tc.SigningKey,
		},
		(&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return token, err
	}

	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{
			Algorithm: jose.DIRECT,
			Key:       tc.EncryptionKey,
		},
		(&jose.EncrypterOptions{}).WithType("JWT").WithContentType("JWT"))
	if err != nil {
		return token, err
	}

	token, err = jwt.SignedAndEncrypted(signer, encrypter).Claims(claims).CompactSerialize()

	return token, err
}

// Takes the encrypted token and a pointer to the claims, the token is verified and deserialized into the provided claims pointer.
func (tc *tokenizer) GetClaims(ctx context.Context, secureToken string, claims Validateable) error {

	if reflect.ValueOf(claims).Kind() != reflect.Pointer {
		return ErrNonPointerClaim
	}

	parsedToken, err := jwt.ParseSignedAndEncrypted(secureToken)
	if err != nil {
		return err
	}

	nested, err := parsedToken.Decrypt(tc.EncryptionKey)
	if err != nil {
		return err
	}

	if err := nested.Claims(tc.SigningKey, claims); err != nil {
		return err
	}

	return claims.Validate()
}

// Takes the encrypted token and a pointer to the claims, the token is deserialized into the provided claims pointer without verification.
func (tc *tokenizer) GetClaimsUnsafe(ctx context.Context, secureToken string, claims Validateable) error {

	if reflect.ValueOf(claims).Kind() != reflect.Pointer {
		return ErrNonPointerClaim
	}

	parsedToken, err := jwt.ParseSignedAndEncrypted(secureToken)
	if err != nil {
		return err
	}

	nested, err := parsedToken.Decrypt(tc.EncryptionKey)
	if err != nil {
		return err
	}

	return nested.UnsafeClaimsWithoutVerification(&claims)
}
