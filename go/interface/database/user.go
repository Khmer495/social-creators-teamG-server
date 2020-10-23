package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) FindById(id int) (user model.User, err error) {
	if err = repo.First(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users model.Users, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) Store(u model.User) (user model.User, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) Update(u model.User) (user model.User, err error) {
	if err = repo.Model(&model.User{}).Update(&u).First(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (repo *UserRepository) DeleteById(user model.User) (err error) {
	if err = repo.Delete(&user).Error; err != nil {
		return
	}
	return
}
