package service

import (
	"context"
	"github.com/Minsoo-Shin/go-boilerplate/domain"
	eu "github.com/Minsoo-Shin/go-boilerplate/internal/user/error"
	mock_repository "github.com/Minsoo-Shin/go-boilerplate/internal/user/repository/mock"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func Test_service_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_repository.NewMockRepository(ctrl)

	fakeUser1 := domain.User{}
	err := faker.FakeData(&fakeUser1)
	assert.NoError(t, err)

	fakeUser1.ID = 1
	fakeUser1.Name = "멤버1번"

	fakeUser2 := domain.User{}
	err = faker.FakeData(&fakeUser2)
	assert.NoError(t, err)

	fakeUser2.ID = 2
	fakeUser2.Name = "멤버2번"

	m.
		EXPECT().
		FindAll(context.Background(), domain.UserFindAllParams{IDs: []uint{1, 2}}).
		Return(domain.Users{fakeUser1, fakeUser2}, nil)

	srv := New(m)

	gots, err := srv.FindAll(context.Background(), domain.UserFindAllRequest{
		IDs: []uint{1, 2},
	})
	assert.Equal(t, gots[0].ID, fakeUser1.ID)
	assert.Equal(t, gots[0].Name, fakeUser1.Name)
	assert.Equal(t, gots[0].Birthdate, fakeUser1.Birthdate)

	assert.Equal(t, gots[1].ID, fakeUser2.ID)
	assert.Equal(t, gots[1].Name, fakeUser2.Name)
	assert.Equal(t, gots[1].Birthdate, fakeUser2.Birthdate)
}

func Test_service_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_repository.NewMockRepository(ctrl)

	fakeUser1 := domain.User{}
	err := faker.FakeData(&fakeUser1)
	assert.NoError(t, err)

	fakeUser1.ID = 1
	fakeUser1.Name = "멤버1번"

	m.
		EXPECT().
		FindByID(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, arg uint) (domain.User, error) {
		t.Logf("%v", arg)
		return fakeUser1, nil
	})

	srv := New(m)

	got, err := srv.Find(context.Background(), domain.UserFindRequest{
		ID: 1,
	})
	assert.Equal(t, got.ID, fakeUser1.ID)
	assert.Equal(t, got.Name, fakeUser1.Name)
}

func Test_service_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_repository.NewMockRepository(ctrl)

	fakeUser := domain.User{}
	err := faker.FakeData(&fakeUser)
	assert.NoError(t, err)

	fakeUser.ID = 1
	fakeUser.Name = "멤버1번"

	t.Run("업데이트 성공", func(t *testing.T) {
		changeUser := fakeUser
		changeUser.Name = "이름 변경"
		changeUser.Birthdate = time.Date(1992, 01, 01, 0, 0, 0, 0, time.UTC)
		// 예상되는 mock data를 넣어준다.
		m.
			EXPECT().
			FindByID(context.Background(), fakeUser.ID).Return(fakeUser, nil)

		m.
			EXPECT().
			Update(context.Background(), changeUser).Return(nil)

		srv := New(m)
		// 변경할 값을 인풋값으로 넣어준다.
		err = srv.Update(context.Background(), domain.UserUpdateRequest{
			ID:        changeUser.ID,
			Name:      changeUser.Name,
			Birthdate: changeUser.Birthdate,
		})
		assert.NoError(t, err)
	})

	t.Run("ID가 DB에 없는 경우", func(t *testing.T) {
		changeUser := fakeUser
		changeUser.Name = "이름 변경"
		changeUser.Birthdate = time.Date(1992, 01, 01, 0, 0, 0, 0, time.UTC)
		// 예상되는 mock data를 넣어준다.
		m.
			EXPECT().
			FindByID(context.Background(), fakeUser.ID).Return(domain.User{}, eu.ErrUserNotFound)

		srv := New(m)
		// 변경할 값을 인풋값으로 넣어준다.
		err = srv.Update(context.Background(), domain.UserUpdateRequest{
			ID:        changeUser.ID,
			Name:      changeUser.Name,
			Birthdate: changeUser.Birthdate,
		})
		assert.ErrorIs(t, err, eu.ErrUserNotFound)
	})

}

func Test_service_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_repository.NewMockRepository(ctrl)

	fakeUser := domain.UserCreateRequest{}
	err := faker.FakeData(&fakeUser)
	assert.NoError(t, err)

	m.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, user domain.User) error {
		assert.Equal(t, fakeUser.Name, user.Name)
		assert.Equal(t, fakeUser.Birthdate, user.Birthdate)
		return nil
	}).
		Return(nil)

	srv := New(m)

	err = srv.Create(context.Background(), fakeUser)
	assert.NoErrorf(t, err, "%v", err)
}
