package service

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
)

// 현재는 학생만 고려하여 설계되어있음.
func (s service) Save(ctx context.Context, request domain.UserSaveParams) error {
	return nil
}

func (s service) Find(ctx context.Context, request domain.UserFindParams) (domain.User, error) {
	return domain.User{}, nil
}

func (s service) Update(ctx context.Context, request domain.UpdateUserRequest) error {
	if err := s.repository.Update(ctx, request.ToUserUpdateParams()); err != nil {
		return err
	}
	return nil
}

func (s service) Delete(ctx context.Context, request domain.UserDeleteRequest) error {
	if err := s.repository.Delete(ctx, domain.UserDeleteParams{request.ID}); err != nil {
		return err
	}

	return nil
}

func (s service) FindAll(ctx context.Context, request domain.UserFindAllRequest) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}
