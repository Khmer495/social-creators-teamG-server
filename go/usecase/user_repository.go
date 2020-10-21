package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type UserRepository interface {
	FindById(id int) (model.User, error)
	FindAll() (model.Users, error)
	Store(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	DeleteById(model.User) error
}
