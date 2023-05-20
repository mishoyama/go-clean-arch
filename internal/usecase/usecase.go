package usecase

import (
	"go-clean-arch/internal/repository"
)

type UserUseCase struct {
	userOutput     UserOutput
	userRepository repository.UserRepository
}

func (u UserUseCase) Get(id int) error {
	// todo get user from DB
	user, err := u.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	u.userOutput.Render(*user)

	return nil
}

func NewUserUseCase(userOutput UserOutput, userRepository repository.UserRepository) UserInput {
	return UserUseCase{userOutput, userRepository}
}
