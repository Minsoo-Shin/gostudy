package service

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
	"github.com/Minsoo-Shin/go-boilerplate/internal/user/repository"
)

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

type Service interface {
	Find(ctx context.Context, request domain.GetUserRequest) (domain.GetUserResponse, error)
	Update(ctx context.Context, request domain.UpdateUserRequest) error
	Delete(ctx context.Context, request domain.UserDeleteRequest) error
	FindAll(ctx context.Context, request domain.UserFindAllRequest) error
}
