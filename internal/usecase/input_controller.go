package usecase

type UserInput interface {
	Get(id int) error
}
