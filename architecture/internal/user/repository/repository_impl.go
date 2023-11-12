package repository

import (
	"context"
	"fmt"
	"github.com/Minsoo-Shin/go-boilerplate/entity"
	eu "github.com/Minsoo-Shin/go-boilerplate/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

const (
	_userCollection = "user"
)

func (r repository) Save(ctx context.Context, params entity.UserSaveParams) error {
	if _, err := r.db.Collection(_userCollection).InsertOne(ctx, params.UserInfo); err != nil {
		return eu.InternalError(err)
	}
	return nil
}

func (r repository) Update(ctx context.Context, params entity.UserUpdateParams) error {
	if result := r.db.Collection(_userCollection).FindOneAndUpdate(ctx, params.Filter(), params.Update()); result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return eu.UserError(http.StatusNotFound).WithMessage("Not found")
		}
		return eu.InternalError(result.Err())
	}
	return nil
}

func (r repository) Delete(ctx context.Context, params entity.UserDeleteParams) error {
	result, err := r.db.Collection(_userCollection).DeleteOne(ctx, params.Filter())
	fmt.Println(result.DeletedCount, err)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Find(ctx context.Context, params entity.UserFindParams) (entity.UserInfo, error) {
	var userInfo entity.UserInfo

	err := r.db.Collection(_userCollection).FindOne(ctx, params.Filter()).Decode(&userInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.UserInfo{}, eu.UserError(http.StatusNotFound).WithMessage("user not found")
		}
		return entity.UserInfo{}, eu.InternalError(err)
	}

	return userInfo, nil
}

func (r repository) FindAll(ctx context.Context, params entity.UserFindAllParams) ([]entity.UserInfo, error) {
	cursor, err := r.db.Collection(_userCollection).Find(ctx, bson.M{"matchedTeacher._id": params.TeacherID})
	if err != nil {
		return nil, eu.InternalError(err)
	}

	var userInfos []entity.UserInfo
	if err := cursor.All(ctx, &userInfos); err != nil {
		return nil, eu.InternalError(err)
	}

	return userInfos, nil
}

func (r repository) CheckDuplicatedUserField(ctx context.Context, params entity.UserFindParams) (bool, error) {
	count, err := r.db.Collection(_userCollection).CountDocuments(ctx, params.Filter())
	if err != nil {
		return true, eu.InternalError(err)
	}

	return count != 0, nil
}
