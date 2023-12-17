package domain

import (
	eu "github.com/Minsoo-Shin/go-boilerplate/pkg/errors"
	"time"
)

type UserCreateRequest struct {
	Name      string
	Birthdate time.Time
	Email     string
	Username  string
	Password  string
}

type UserUpdateRequest struct {
	ID        uint
	Name      string
	Birthdate time.Time
}

type UserDeleteRequest struct {
	ID uint
}

type UserFindRequest struct {
	ID uint
}

func (u UserFindRequest) Valid() error {
	if u.ID == 0 {
		return eu.ErrBadRequest
	}
	return nil
}

type UserFindAllRequest struct {
	IDs []uint
}

type UserDto struct {
	ID        uint `gorm:"primarykey" json:"id"`
	Name      string
	Birthdate time.Time
}

type UserDtos []UserDto

func NewUserDtos() UserDtos {
	return UserDtos{}
}

func (u UserDtos) From(users Users) UserDtos {
	dtos := make(UserDtos, 0)
	for _, v := range users {
		dtos = append(dtos, NewUserDto().From(v))
	}
	return dtos
}

func NewUserDto() UserDto {
	return UserDto{}
}

func (u UserDto) From(user User) UserDto {
	return UserDto{
		ID:        user.ID,
		Name:      user.Name,
		Birthdate: user.Birthdate,
	}
}

type UserFindAllParams struct {
	IDs []uint
}
