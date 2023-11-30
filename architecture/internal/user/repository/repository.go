package repository

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

type Repository interface {
	Create(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User) error
	FindByID(ctx context.Context, userID uint) (domain.User, error)
	FindAll(ctx context.Context)
	Delete(ctx context.Context, userID uint) error
}
