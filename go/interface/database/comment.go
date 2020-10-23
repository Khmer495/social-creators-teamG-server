package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type CommentRepository struct {
	SqlHandler
}

func (repo *CommentRepository) FindById(id int) (comment model.Comment, err error) {
	if err = repo.Set("gorm:auto_preload", true).First(&comment, id).Error; err != nil {
		return
	}
	return
}

func (repo *CommentRepository) FindAll(m model.CommentQuery) (comments model.Comments, err error) {
	query := repo.Set("gorm:auto_preload", true)
	if m.UserID > 0 {
		query = query.Where("user_id = ?", m.UserID)
	}
	if m.RestaurantID > 0 {
		query = query.Where("restaurant_id = ?", m.RestaurantID)
	}
	if err = query.Find(&comments).Error; err != nil {
		return
	}
	return
}

func (repo *CommentRepository) Store(c model.Comment) (perfecture model.Comment, err error) {
	if err = repo.Set("gorm:auto_preload", true).Create(&c).Error; err != nil {
		return
	}
	perfecture = c
	return
}

func (repo *CommentRepository) Update(c model.Comment) (perfecture model.Comment, err error) {
	if err = repo.Set("gorm:auto_preload", true).Model(&model.Comment{}).Update(&c).First(&c).Error; err != nil {
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
