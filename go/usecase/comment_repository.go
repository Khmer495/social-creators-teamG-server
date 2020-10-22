package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type CommentRepository interface {
	FindById(id int) (model.Comment, error)
	FindAll() (model.Comments, error)
	Store(model.Comment) (model.Comment, error)
	Update(model.Comment) (model.Comment, error)
	DeleteById(model.Comment) error
}
