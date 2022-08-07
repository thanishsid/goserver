package directive

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/thanishsid/goserver/domain"
)

var ErrUnauthorized = errors.New("you are not authorized to perform this action")

func Authorize(ctx context.Context, obj interface{}, next graphql.Resolver, rolesStrPtr *string, allowOwnerPtr, mustOwnPtr *bool) (res interface{}, err error) {
	checkRoles := rolesStrPtr != nil && *rolesStrPtr != ""
	checkAllowOwner := allowOwnerPtr != nil && *allowOwnerPtr == true
	checkMustOwn := mustOwnPtr != nil && *mustOwnPtr == true

	authorized := false

	session, err := domain.SessionFor(ctx)
	if err != nil {
		return nil, ErrUnauthorized
	}

	if checkRoles {
		roles, err := getRolesFromString(*rolesStrPtr)
		if err != nil {
			return nil, err
		}

		for _, role := range roles {
			if role == session.UserRole {
				authorized = true
				break
			}
		}
	}

	if checkAllowOwner || checkMustOwn {
		if obj == nil {
			return nil, errors.New("cannot determine owner on a null object")
		}

		isOwner := false

		// Assert different entity types and check if current session user owns the entity.
		switch t := obj.(type) {
		case *domain.User:
			isOwner = t.IsOwner(session.UserID)
		}

		if checkMustOwn && !isOwner {
			return nil, ErrUnauthorized
		}

		if isOwner {
			authorized = true
		}
	}

	if authorized {
		return next(ctx)
	}

	return nil, ErrUnauthorized
}

// Get roles from schema directive.
func getRolesFromString(rolesStr string) ([]domain.Role, error) {
	rolesStrings := strings.Split(rolesStr, ",")

	roles := make([]domain.Role, len(rolesStrings))

	for i, roleStr := range rolesStrings {
		role := domain.Role(strings.Trim(roleStr, " "))

		if err := role.ValidateRole(); err != nil {
			return nil, fmt.Errorf("invalid role, '%s' passed to 'hasRole' schema directive.", role)
		}

		roles[i] = role
	}

	return roles, nil
}
