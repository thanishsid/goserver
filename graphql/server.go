package graphql

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/cookiemanager"
	"github.com/thanishsid/goserver/graphql/dataloader"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/resolver"
)

type ServerDeps struct {
	UserService    domain.UserService
	ImageService   domain.ImageService
	SessionService domain.SessionService
}

func NewServer(deps ServerDeps) *chi.Mux {

	// Create new mux router.
	mux := chi.NewRouter()

	cookieConfig := cookiemanager.CookieConfig{
		HttpOnly: true,
		Secure:   config.C.Environment != "development",
		Path:     "/",
		Domain:   "127.0.0.1",
	}

	// Load middleware.
	mux.Use(
		cookiemanager.LoadManager(cookieConfig),
		LoadSessionMiddleware(deps.SessionService),
		dataloader.Middleware(dataloader.NewDataloader(deps.UserService, deps.ImageService)),
	)

	// Graphql configs.
	conf := generated.Config{Resolvers: &resolver.Resolver{
		UserService:    deps.UserService,
		ImageService:   deps.ImageService,
		SessionService: deps.SessionService,
	},
		Directives: generated.DirectiveRoot{
			EnforceAction: enforceActionDirective,
		},
	}

	// Create graphqQL server.
	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(conf))

	// Graphql Error handler.
	gqlHandler.SetErrorPresenter(gqlErrFunc)

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", gqlHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.C.ServerPort)

	return mux
}

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
