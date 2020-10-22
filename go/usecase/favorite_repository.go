package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type FavoriteRepository interface {
	FindById(id int) (model.Favorite, error)
	FindAll(model.FavoriteQuery) (model.Favorites, error)
	Store(model.Favorite) (model.Favorite, error)
	Update(model.Favorite) (model.Favorite, error)
	DeleteById(model.Favorite) error
}
