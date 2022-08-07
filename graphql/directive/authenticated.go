package directive

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/thanishsid/goserver/domain"
)

func Authenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	_, err = domain.SessionFor(ctx)
	if err != nil {
		return nil, errors.New("you need to be logged in to perform this action")
	}

	return next(ctx)
}
