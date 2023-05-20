package presenter

import (
	"context"
	"go-clean-arch/internal/model"
)

type User struct {
	ctx context.Context
}

func NewUser(ctx context.Context) User {
	return User{ctx}
}

func (u User) Render(user model.User) {
	u.ctx.Value(user)
}
