package service

import (
	"context"
	"gostudy/architecture/entity"
	"gostudy/architecture/internal/user/repository"
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
	Find(ctx context.Context, request entity.GetUserRequest) (entity.GetUserResponse, error)
	Update(ctx context.Context, request entity.UpdateUserRequest) error
	Delete(ctx context.Context, request entity.UserDeleteRequest) error
	FindAll(ctx context.Context, request entity.UserFindAllRequest) error
}
