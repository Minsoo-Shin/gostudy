package service

import (
	"context"
	"ggurugi/entity"
)

// 현재는 학생만 고려하여 설계되어있음.
func (s service) Save(ctx context.Context, request entity.UserSaveParams) error {
	return nil
}

func (s service) Find(ctx context.Context, request entity.UserFindParams) (entity.User, error) {
	return entity.User{}, nil
}

func (s service) Update(ctx context.Context, request entity.UpdateUserRequest) error {
	if err := s.repository.Update(ctx, request.ToUserUpdateParams()); err != nil {
		return err
	}
	return nil
}

func (s service) Delete(ctx context.Context, request entity.UserDeleteRequest) error {
	if err := s.repository.Delete(ctx, entity.UserDeleteParams{request.ID}); err != nil {
		return err
	}

	return nil
}

func (s service) FindAll(ctx context.Context, request entity.UserFindAllRequest) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
