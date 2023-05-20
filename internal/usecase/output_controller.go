package usecase

import "go-clean-arch/internal/model"

type UserOutput interface {
	Render(user model.User)
}
