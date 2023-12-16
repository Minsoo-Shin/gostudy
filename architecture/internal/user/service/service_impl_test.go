package service

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
	mock_repository "github.com/Minsoo-Shin/go-boilerplate/internal/user/repository/mock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_service_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_repository.NewMockRepository(ctrl)

	fakeUser := domain.User{}
	err := faker.FakeData(&fakeUser)
	assert.NoError(t, err)

	fakeUser.ID = 1
	fakeUser.Name = "멤버1번"

	m.
		EXPECT().
		FindAll(context.Background(), domain.UserFindAllParams{IDs: []uint{1}}).
		Return(domain.Users{fakeUser}, nil)

	srv := New(m)

	gots, err := srv.FindAll(context.Background(), domain.UserFindAllRequest{
		IDs: []uint{1},
	})
	assert.Equal(t, gots[0].Name, fakeUser.Name)
}
