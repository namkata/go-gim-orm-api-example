package userbiz

import (
	"context"
	"errors"
	usermodel "local-app/modules/users/models"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *usermodel.User) error
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type createUserBiz struct {
	store CreateUserStorage // Store for book creation
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *usermodel.User) error {
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return errors.New("User already exists!")
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
