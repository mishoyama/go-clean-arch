package gateway

import "go-clean-arch/internal/model"

type Database struct {
	// todo - ORM
}

func (d Database) Find(id int) (*model.User, error) {
	_ = id
	return nil, nil
}

type User struct {
	db Database
}

func NewUser(db Database) User {
	return User{db}
}

func (u User) GetUserById(id int) (*model.User, error) {
	return u.db.Find(id)
}
