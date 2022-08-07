package setup

import (
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
)

type Initializer interface {
	Initialize() error
}

type initializer struct {
	dbs         db.DB
	userService domain.UserService
}

func (i *initializer) Initialize() error {
	return nil
}
