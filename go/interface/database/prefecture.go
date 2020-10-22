package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type PrefectureRepository struct {
	SqlHandler
}

func (repo *PrefectureRepository) FindById(id int) (prefecture model.Prefecture, err error) {
	if err = repo.Find(&prefecture, id).Error; err != nil {
		return
	}
	return
}

func (repo *PrefectureRepository) FindAll() (prefectures model.Prefectures, err error) {
	if err = repo.Find(&prefectures).Error; err != nil {
		return
	}
	return
}

func (repo *PrefectureRepository) Store(p model.Prefecture) (perfecture model.Prefecture, err error) {
	if err = repo.Create(&p).Error; err != nil {
		return
	}
	perfecture = p
	return
}

func (repo *PrefectureRepository) Update(p model.Prefecture) (perfecture model.Prefecture, err error) {
	if err = repo.Save(&p).Error; err != nil {
		return
	}
	perfecture = p
	return
}

func (repo *PrefectureRepository) DeleteById(prefecture model.Prefecture) (err error) {
	if err = repo.Delete(&prefecture).Error; err != nil {
		return
	}
	return
}
