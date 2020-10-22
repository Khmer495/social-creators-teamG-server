package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type PrefectureRepository interface {
	FindById(id int) (model.Prefecture, error)
	FindAll() (model.Prefectures, error)
	Store(model.Prefecture) (model.Prefecture, error)
	Update(model.Prefecture) (model.Prefecture, error)
	DeleteById(model.Prefecture) error
}
