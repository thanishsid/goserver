package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func gqlErrFunc(ctx context.Context, err error) *gqlerror.Error {
	gqlErr := graphql.DefaultErrorPresenter(ctx, err)

	if gqlErr.Extensions == nil {
		gqlErr.Extensions = make(map[string]interface{})
	}

	var vdErr validation.Errors
	if errors.As(err, &vdErr) {
		gqlErr.Message = "validation error"
		gqlErr.Extensions["validationErrors"] = vdErr
	}

	return gqlErr
}
