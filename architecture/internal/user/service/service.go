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
	Create(ctx context.Context, request domain.UserCreateRequest) error
	Find(ctx context.Context, request domain.UserFindRequest) (domain.UserDto, error)
	Update(ctx context.Context, request domain.UserUpdateRequest) error
	Delete(ctx context.Context, request domain.UserDeleteRequest) error
	FindAll(ctx context.Context, request domain.UserFindAllRequest) (domain.UserDtos, error)
}
