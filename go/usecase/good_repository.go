package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type GoodRepository interface {
	FindById(id int) (model.Good, error)
	FindAll() (model.Goods, error)
	Store(model.Good) (model.Good, error)
	Update(model.Good) (model.Good, error)
	DeleteById(model.Good) error
}
