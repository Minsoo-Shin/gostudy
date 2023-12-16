package service

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
)

func (s service) Find(ctx context.Context, request domain.UserFindRequest) (domain.UserDto, error) {
	return domain.NewUserDto(), nil
}

func (s service) Update(ctx context.Context, request domain.UserUpdateRequest) error {
	if err := s.repository.Update(ctx, domain.User{}); err != nil {
		return err
	}
	return nil
}

func (s service) Delete(ctx context.Context, request domain.UserDeleteRequest) error {
	if err := s.repository.Delete(ctx, request.ID); err != nil {
		return err
	}

	return nil
}

func (s service) FindAll(ctx context.Context, request domain.UserFindAllRequest) (domain.UserDtos, error) {
	users, err := s.repository.FindAll(ctx, domain.UserFindAllParams{
		IDs: request.IDs,
	})
	if err != nil {
		return nil, err
	}

	return domain.NewUserDtos().From(users), nil
}
