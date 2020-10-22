package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type FavoriteRepository struct {
	SqlHandler
}

func (repo *FavoriteRepository) FindById(id int) (favorite model.Favorite, err error) {
	if err = repo.Find(&favorite, id).Error; err != nil {
		return
	}
	return
}

func (repo *FavoriteRepository) FindAll() (favorites model.Favorites, err error) {
	if err = repo.Find(&favorites).Error; err != nil {
		return
	}
	return
}

func (repo *FavoriteRepository) Store(f model.Favorite) (perfecture model.Favorite, err error) {
	if err = repo.Create(&f).Error; err != nil {
		return
	}
	perfecture = f
	return
}

func (repo *FavoriteRepository) Update(f model.Favorite) (perfecture model.Favorite, err error) {
	if err = repo.Save(&f).Error; err != nil {
		return
	}
	perfecture = f
	return
}

func (repo *FavoriteRepository) DeleteById(favorite model.Favorite) (err error) {
	if err = repo.Delete(&favorite).Error; err != nil {
		return
	}
	return
}
