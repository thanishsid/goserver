package security

import (
	"context"
	"io/fs"

	"github.com/99designs/gqlgen/graphql"
	"github.com/casbin/casbin/v2"
	casbin_fs_adapter "github.com/naucon/casbin-fs-adapter"

	"github.com/thanishsid/goserver/domain"
)

func NewAuthorizer(fsys fs.FS) (*Authorizer, error) {
	model, err := casbin_fs_adapter.NewModel(fsys, "casbin-model.conf")
	if err != nil {
		return nil, err
	}

	policies := casbin_fs_adapter.NewAdapter(fsys, "casbin-policy.csv")

	enforcer, err := casbin.NewEnforcer(model, policies)
	if err != nil {
		return nil, err
	}

	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return &Authorizer{enforcer}, nil
}

type Authorizer struct {
	*casbin.Enforcer
}

func (a *Authorizer) AuthorizeDirective(ctx context.Context, obj interface{}, next graphql.Resolver, object, action string) (res interface{}, err error) {
	authorized := false

	role := "guest"

	guestIsValid, err := a.Enforce(role, object, action)
	if err != nil {
		return nil, err
	}

	session, err := domain.SessionFor(ctx)
	if err == nil {
		role = session.UserRole.String()
	}

	authorized, err = a.Enforce(role, object, action)
	if err != nil {
		return nil, err
	}

	if !authorized && guestIsValid {
		return nil, ErrMustBeLoggedOut
	}

	if !authorized {
		authorized = checkObjOwnership(obj, session)
	}

	if authorized {
		return next(ctx)
	}

	if role == "guest" && !authorized {
		return nil, ErrNotLoggedIn
	}

	return nil, ErrUnauthorized
}

// Check ownership of an object passed to graphql schema directive.
func checkObjOwnership(obj any, session *domain.Session) bool {
	if session == nil {
		return false
	}

	switch t := obj.(type) {
	case *domain.User:
		return t.IsOwner(session.UserID)
	}

	return false
}
