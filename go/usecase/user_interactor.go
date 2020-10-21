package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UserById(id int) (user model.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) Users() (users model.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) Add(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) DeleteById(u model.User) (err error) {
	err = interactor.UserRepository.DeleteById(u)
	return
}
