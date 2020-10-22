package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type CommentRepository struct {
	SqlHandler
}

func (repo *CommentRepository) FindById(id int) (comment model.Comment, err error) {
	if err = repo.Find(&comment, id).Error; err != nil {
		return
	}
	return
}

func (repo *CommentRepository) FindAll() (comments model.Comments, err error) {
	if err = repo.Find(&comments).Error; err != nil {
		return
	}
	return
}

func (repo *CommentRepository) Store(c model.Comment) (perfecture model.Comment, err error) {
	if err = repo.Create(&c).Error; err != nil {
		return
	}
	perfecture = c
	return
}

func (repo *CommentRepository) Update(c model.Comment) (perfecture model.Comment, err error) {
	if err = repo.Save(&c).Error; err != nil {
		return
	}
	perfecture = c
	return
}

func (repo *CommentRepository) DeleteById(comment model.Comment) (err error) {
	if err = repo.Delete(&comment).Error; err != nil {
		return
	}
	return
}
