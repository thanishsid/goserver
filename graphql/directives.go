package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

const GRAPHQL_ACTION_CTX_KEY = "graphql-action"

func enforceActionDirective(ctx context.Context, obj interface{}, next graphql.Resolver, action string) (res interface{}, err error) {
	ctxWithAction := context.WithValue(ctx, GRAPHQL_ACTION_CTX_KEY, action)
	// fmt.Printf("%+v\n", graphql.GetFieldContext(ctx).Args)
	return next(ctxWithAction)
}
