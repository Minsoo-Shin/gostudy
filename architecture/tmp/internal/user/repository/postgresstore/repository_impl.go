package postgresstore

import (
	"context"
	"fmt"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
	eu "github.com/Minsoo-Shin/go-boilerplate/internal/user/error"
	"github.com/Minsoo-Shin/go-boilerplate/internal/user/repository"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type userPostgresRepo struct {
	db *gorm.DB
}

func (r userPostgresRepo) Create(ctx context.Context, user domain.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return fmt.Errorf("%w", eu.ErrNotDefined)
	}
	return nil
}

func (r userPostgresRepo) Update(ctx context.Context, user domain.User) error {
	if err := r.db.Save(&user).Error; err != nil {
		return eu.ErrNotDefined
	}
	return nil
}

func (r userPostgresRepo) FindByID(ctx context.Context, userID uint) (domain.User, error) {
	var user = domain.User{ID: userID}

	if err := r.db.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("%w", eu.ErrUserNotFound)
		}
		return user, fmt.Errorf("%w", eu.ErrNotDefined)
	}
	return user, nil
}

func (r userPostgresRepo) FindAll(ctx context.Context, params domain.UserFindAllParams) (domain.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (r userPostgresRepo) Delete(ctx context.Context, userID uint) error {
	//TODO implement me
	panic("implement me")
}

func New(db *gorm.DB) repository.Repository {
	return &userPostgresRepo{
		db: db,
	}
}
