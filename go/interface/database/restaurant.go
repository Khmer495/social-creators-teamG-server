package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type RestaurantRepository struct {
	SqlHandler
}

func (repo *RestaurantRepository) FindById(id int) (restaurant model.Restaurant, err error) {
	if err = repo.Find(&restaurant, id).Error; err != nil {
		return
	}
	return
}

func (repo *RestaurantRepository) FindAll() (restaurants model.Restaurants, err error) {
	if err = repo.Find(&restaurants).Error; err != nil {
		return
	}
	return
}

func (repo *RestaurantRepository) Store(r model.Restaurant) (perfecture model.Restaurant, err error) {
	if err = repo.Create(&r).Error; err != nil {
		return
	}
	perfecture = r
	return
}

func (repo *RestaurantRepository) Update(r model.Restaurant) (perfecture model.Restaurant, err error) {
	if err = repo.Save(&r).Error; err != nil {
		return
	}
	perfecture = r
	return
}

func (repo *RestaurantRepository) DeleteById(restaurant model.Restaurant) (err error) {
	if err = repo.Delete(&restaurant).Error; err != nil {
		return
	}
	return
}
