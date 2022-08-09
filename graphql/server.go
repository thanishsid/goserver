package graphql

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/thanishsid/goserver/config"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/graphql/cookiemanager"
	"github.com/thanishsid/goserver/graphql/dataloader"
	"github.com/thanishsid/goserver/graphql/generated"
	"github.com/thanishsid/goserver/graphql/middleware"
	"github.com/thanishsid/goserver/graphql/resolver"
	"github.com/thanishsid/goserver/infrastructure/security"
)

type ServerDeps struct {
	UserService    domain.UserService
	ImageService   domain.ImageService
	SessionService domain.SessionService
	AuthService    domain.AuthService
	Authorizer     *security.Authorizer
}

func NewServer(deps ServerDeps) *chi.Mux {

	// Create new mux router.
	r := chi.NewRouter()

	// Cookie Manager Configs.
	cookieConfig := cookiemanager.CookieConfig{
		HttpOnly: true,
		Secure:   config.C.Environment != "development",
		Path:     "/",
	}

	// Load middleware.
	r.Use(
		middleware.LoadRequestMetadataMiddleware(),
		middleware.LoadCookieManager(cookieConfig),
		middleware.LoadSessionMiddleware(deps.SessionService),
		dataloader.Middleware(dataloader.NewDataloader(deps.UserService, deps.ImageService)),
	)

	// Mount the Oauth2 login http routes.
	r.Mount("/auth", LoadOauthRoutes(deps.AuthService))

	// Load graphql handler onto http mux in the /query endpoint
	r.Handle("/query", NewGraphqlHandler(deps))

	// Will load GraphiQL if in Development Environment
	if config.C.Environment == "development" {
		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.C.ServerPort)
	}

	return r
}

// Create a new graphql handler with configs.
func NewGraphqlHandler(deps ServerDeps) *handler.Server {
	// Graphql configs.
	conf := generated.Config{Resolvers: &resolver.Resolver{
		UserService:    deps.UserService,
		ImageService:   deps.ImageService,
		SessionService: deps.SessionService,
		AuthService:    deps.AuthService,
	},
		Directives: generated.DirectiveRoot{
			Authorize: deps.Authorizer.AuthorizeDirective,
		},
	}

	// Create graphqQL server.
	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(conf))

	// Graphql Error handler.
	gqlHandler.SetErrorPresenter(gqlErrFunc)

	return gqlHandler
}
