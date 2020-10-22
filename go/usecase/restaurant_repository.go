package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type RestaurantRepository interface {
	FindById(id int) (model.Restaurant, error)
	FindAll() (model.Restaurants, error)
	Store(model.Restaurant) (model.Restaurant, error)
	Update(model.Restaurant) (model.Restaurant, error)
	DeleteById(model.Restaurant) error
}
