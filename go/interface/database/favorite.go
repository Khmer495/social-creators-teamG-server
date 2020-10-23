package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type FavoriteRepository struct {
	SqlHandler
}

func (repo *FavoriteRepository) FindById(id int) (favorite model.Favorite, err error) {
	if err = repo.Set("gorm:auto_preload", true).First(&favorite, id).Error; err != nil {
		return
	}
	return
}

func (repo *FavoriteRepository) FindAll(f model.FavoriteQuery) (favorites model.Favorites, err error) {
	query := repo.Set("gorm:auto_preload", true)
	if f.UserID > 0 {
		query = query.Where("user_id = ?", f.UserID)
	}
	if f.RestaurantID > 0 {
		query = query.Where("restaurant_id = ?", f.RestaurantID)
	}
	if err = query.Find(&favorites).Error; err != nil {
		return
	}
	return
}

func (repo *FavoriteRepository) Store(f model.Favorite) (perfecture model.Favorite, err error) {
	if err = repo.Set("gorm:auto_preload", true).Create(&f).Error; err != nil {
		return
	}
	perfecture = f
	return
}

func (repo *FavoriteRepository) Update(f model.Favorite) (perfecture model.Favorite, err error) {
	if err = repo.Set("gorm:auto_preload", true).Model(&model.Favorite{}).Update(&f).First(&f).Error; err != nil {
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
