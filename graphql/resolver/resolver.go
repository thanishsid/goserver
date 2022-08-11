package resolver

import (
	"github.com/thanishsid/goserver/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService    domain.UserService
	ImageService   domain.ImageService
	VideoService   domain.VideoService
	SessionService domain.SessionService
	AuthService    domain.AuthService
}
