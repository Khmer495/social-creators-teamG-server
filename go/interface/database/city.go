package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type CityRepository struct {
	SqlHandler
}

func (repo *CityRepository) FindById(id int) (city model.City, err error) {
	if err = repo.Set("gorm:auto_preload", true).First(&city, id).Error; err != nil {
		return
	}
	return
}

func (repo *CityRepository) FindAll() (citys model.Cities, err error) {
	if err = repo.Set("gorm:auto_preload", true).Find(&citys).Error; err != nil {
		return
	}
	return
}

func (repo *CityRepository) Store(c model.City) (perfecture model.City, err error) {
	if err = repo.Set("gorm:auto_preload", true).Create(&c).Error; err != nil {
		return
	}
	perfecture = c
	return
}

func (repo *CityRepository) Update(c model.City) (perfecture model.City, err error) {
	if err = repo.Set("gorm:auto_preload", true).Model(&model.City{}).Update(&c).First(&c).Error; err != nil {
		return
	}
	perfecture = c
	return
}

func (repo *CityRepository) DeleteById(city model.City) (err error) {
	if err = repo.Delete(&city).Error; err != nil {
		return
	}
	return
}
