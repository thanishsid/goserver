package service

import (
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/internal/mailer"
	"github.com/thanishsid/goserver/internal/tokenizer"
	"github.com/thanishsid/goserver/repository"
)

// Inititate and get all services.
func New(deps *ServiceDeps) Service {
	return &svc{
		userService:  &userService{deps},
		imageService: &imageService{deps},
	}
}

type ServiceDeps struct {
	Tokens tokenizer.Tokenizer
	Mail   mailer.Mailer
	Repo   repository.Repository
}

type Service interface {
	UserService() domain.UserService
	ImageService() domain.ImageService
}

type svc struct {
	userService  domain.UserService
	imageService domain.ImageService
}

func (s *svc) UserService() domain.UserService {
	return s.userService
}

func (s *svc) ImageService() domain.ImageService {
	return s.imageService
}
