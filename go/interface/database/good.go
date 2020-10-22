package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type GoodRepository struct {
	SqlHandler
}

func (repo *GoodRepository) FindById(id int) (good model.Good, err error) {
	if err = repo.Find(&good, id).Error; err != nil {
		return
	}
	return
}

func (repo *GoodRepository) FindAll() (goods model.Goods, err error) {
	if err = repo.Find(&goods).Error; err != nil {
		return
	}
	return
}

func (repo *GoodRepository) Store(g model.Good) (perfecture model.Good, err error) {
	if err = repo.Create(&g).Error; err != nil {
		return
	}
	perfecture = g
	return
}

func (repo *GoodRepository) Update(g model.Good) (perfecture model.Good, err error) {
	if err = repo.Save(&g).Error; err != nil {
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
