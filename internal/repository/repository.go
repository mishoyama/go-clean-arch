package repository

import "go-clean-arch/internal/model"

type UserRepository interface {
	GetUserById(id int) (*model.User, error)
}
