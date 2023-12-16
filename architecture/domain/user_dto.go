package domain

import (
	"time"
)

type UserUpdateRequest struct {
}

type UserDeleteRequest struct {
	ID uint
}

type UserFindRequest struct {
	ID uint
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
