package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type GoodRepository struct {
	SqlHandler
}

func (repo *GoodRepository) FindById(id int) (good model.Good, err error) {
	if err = repo.Set("gorm:auto_preload", true).First(&good, id).Error; err != nil {
		return
	}
	return
}

func (repo *GoodRepository) FindAll(g model.GoodQuery) (goods model.Goods, err error) {
	query := repo.Set("gorm:auto_preload", true)
	if g.UserID > 0 {
		query = query.Where("user_id = ?", g.UserID)
	}
	if g.RestaurantID > 0 {
		query = query.Where("restaurant_id = ?", g.RestaurantID)
	}
	if err = query.Find(&goods).Error; err != nil {
		return
	}
	return
}

func (repo *GoodRepository) Store(g model.Good) (perfecture model.Good, err error) {
	if err = repo.Set("gorm:auto_preload", true).Create(&g).Error; err != nil {
		return
	}
	perfecture = g
	return
}

func (repo *GoodRepository) Update(g model.Good) (perfecture model.Good, err error) {
	if err = repo.Set("gorm:auto_preload", true).Model(&model.Good{}).Update(&g).First(&g).Error; err != nil {
		return
	}
	perfecture = g
	return
}

func (repo *GoodRepository) DeleteById(good model.Good) (err error) {
	if err = repo.Delete(&good).Error; err != nil {
		return
	}
	return
}
