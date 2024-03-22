package main

import (
	"fmt"

	"go.uber.org/fx"
)

func main() {
	// Fx 애플리케이션을 시작합니다.
	app := fx.New(
		// 여러 모듈을 제공합니다.
		fx.Provide(
			provideService("Service"),
		),
		// 모든 모듈을 시작합니다.
		fx.Invoke(run),
	)

	// 애플리케이션을 시작하고 실행합니다.
	app.Run()
}

// 서비스 생성 함수
func provideService(name string) func() []*Service {
	return func() []*Service {
		var services []*Service
		for i := 0; i < 3; i++ {
			services = append(services, &Service{Name: fmt.Sprintf("%s-%d", name, i+1)})
		}
		return services
	}
}

// 서비스 타입
type Service struct {
	Name string
}

// 실행 함수
func run(services []*Service) {
	fmt.Println("Running services:")
	for _, service := range services {
		fmt.Println(service.Name)
	}
}
