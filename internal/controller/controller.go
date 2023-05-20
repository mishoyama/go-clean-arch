package controller

import (
	"context"
	"go-clean-arch/internal/gateway"
	"go-clean-arch/internal/presenter"

	"go-clean-arch/internal/usecase"
)

type User struct{}

func (u *User) GetUser(ctx context.Context) error {
	id := 0
	// Pass a specific implementation of Output port and generate the usecase.
	outputPort := presenter.NewUser(ctx)
	repoPort := gateway.NewUser(gateway.Database{})
	userUsecase := usecase.NewUserUseCase(outputPort, repoPort)

	return userUsecase.Get(id)
}
