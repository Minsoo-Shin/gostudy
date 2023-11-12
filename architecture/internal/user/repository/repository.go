package repository

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/entity"
	"github.com/Minsoo-Shin/go-boilerplate/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func New(cfg config.Config, client *mongo.Client) Repository {
	return &repository{
		db: client.Database(cfg.DbName),
	}
}

type Repository interface {
	Save(ctx context.Context, Params entity.UserSaveParams) error
	Find(ctx context.Context, params entity.UserFindParams) (entity.UserInfo, error)
	Update(ctx context.Context, params entity.UserUpdateParams) error
	Delete(ctx context.Context, params entity.UserDeleteParams) error
	FindAll(ctx context.Context, params entity.UserFindAllParams) ([]entity.UserInfo, error)
	CheckDuplicatedUserField(ctx context.Context, params entity.UserFindParams) (bool, error)
}
