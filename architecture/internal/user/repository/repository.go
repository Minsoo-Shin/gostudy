package repository

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
)

type Repository interface {
	Create(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User) error
	FindByID(ctx context.Context, userID uint) (domain.User, error)
	FindAll(ctx context.Context, params domain.UserFindAllParams) (domain.Users, error)
	Delete(ctx context.Context, userID uint) error
}
