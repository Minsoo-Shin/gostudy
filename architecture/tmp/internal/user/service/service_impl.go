package service

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
)

func (s service) Create(ctx context.Context, request domain.UserCreateRequest) error {
	if err := s.repository.Create(ctx, domain.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		Name:      request.Name,
		Birthdate: request.Birthdate,
	}); err != nil {
		return err
	}

	return nil
}

func (s service) Find(ctx context.Context, request domain.UserFindRequest) (domain.UserDto, error) {
	user, err := s.repository.FindByID(ctx, request.ID)
	if err != nil {
		return domain.UserDto{}, err
	}

	return domain.NewUserDto().From(user), nil
}

func (s service) Update(ctx context.Context, request domain.UserUpdateRequest) error {
	user, err := s.repository.FindByID(ctx, request.ID)
	if err != nil {
		return err
	}

	user.Name = request.Name
	user.Birthdate = request.Birthdate

	if err := s.repository.Update(ctx, user); err != nil {
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
