package dataloader

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"
	"github.com/thanishsid/goserver/domain"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Dataloader struct {
	userLoader  *dataloader.Loader
	imageLoader *dataloader.Loader
}

// Instantiate and return new Dataloader.
func NewDataloader(us domain.UserService, is domain.ImageService) *Dataloader {
	users := &userBatcher{us}
	images := &imageBatcher{is}

	return &Dataloader{
		userLoader:  dataloader.NewBatchedLoader(users.Get),
		imageLoader: dataloader.NewBatchedLoader(images.Get),
	}
}

// Inject Dataloader into the middleware stack.
func Middleware(loader *Dataloader) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			nextCtx := context.WithValue(r.Context(), loadersKey, loader)
			next.ServeHTTP(w, r.WithContext(nextCtx))
		})
	}
}

// For returns the dataloader for a given context.
func For(ctx context.Context) *Dataloader {
	return ctx.Value(loadersKey).(*Dataloader)
}

//*-------- Utility Funcs to obtain objects ------------>

func (d *Dataloader) GetUser(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	thunk := d.userLoader.Load(ctx, dataloader.StringKey(userID.String()))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*domain.User), nil
}

func (d *Dataloader) GetImage(ctx context.Context, imageID uuid.UUID) (*domain.Image, error) {
	thunk := d.imageLoader.Load(ctx, dataloader.StringKey(imageID.String()))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*domain.Image), nil
}
