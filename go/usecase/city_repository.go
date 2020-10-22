package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type CityRepository interface {
	FindById(id int) (model.City, error)
	FindAll() (model.Cities, error)
	Store(model.City) (model.City, error)
	Update(model.City) (model.City, error)
	DeleteById(model.City) error
}
