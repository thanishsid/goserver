package service

import (
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/internal/mailer"
	"github.com/thanishsid/goserver/internal/tokenizer"
	"github.com/thanishsid/goserver/repository"
)

// Inititate and get all services.
func New(deps *ServiceDeps) Service {
	return &srvc{
		userService: &userService{deps},
	}
}

type ServiceDeps struct {
	Tokens tokenizer.Tokenizer
	Mail   mailer.Mailer
	Repo   repository.Repository
}

type Service interface {
	UserService() domain.UserService
}

type srvc struct {
	userService domain.UserService
}

func (s *srvc) UserService() domain.UserService {
	return s.userService
}
