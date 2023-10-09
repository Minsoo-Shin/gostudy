package fx

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"testing"
)

var (
	myServiceAddress *MyService
)

type MyService struct{}

func NewMyService(t *testing.T) *MyService {
	t.Run("myServiceAddress는 한번만 실행되며, 동일한 주소로 사용된다.", func(t *testing.T) {
		assert.Equal(t, (*MyService)(nil), myServiceAddress)
	})

	myServiceAddress = &MyService{}
	return myServiceAddress
}

func DoSomethingByMyService1(t *testing.T, myService *MyService) {
	//fmt.Printf("DoSomethingByMyService1: %p\n", myService)
	t.Run("처음에 선언한 NewMyService 메모리 주소로 사용", func(t *testing.T) {
		assert.Equal(t, myServiceAddress, myService)
	})
}

func DoSomethingByMyService2(t *testing.T, myService *MyService) {
	//fmt.Printf("DoSomethingByMyService2: %p\n", myService)
	t.Run("처음에 선언한 NewMyService 메모리 주소로 사용", func(t *testing.T) {
		assert.Equal(t, myServiceAddress, myService)
	})
}

func TestSingletonService(t *testing.T) {
	app := fx.New(
		fx.Provide(
			func() *testing.T { return t },
			NewMyService,
		),
		fx.Invoke(
			DoSomethingByMyService1,
			DoSomethingByMyService2,
		),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())

}
