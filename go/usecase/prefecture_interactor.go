package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type PrefectureInteractor struct {
	PrefectureRepository PrefectureRepository
}

func (interactor *PrefectureInteractor) PrefectureById(id int) (prefecture model.Prefecture, err error) {
	prefecture, err = interactor.PrefectureRepository.FindById(id)
	return
}

func (interactor *PrefectureInteractor) Prefectures() (prefectures model.Prefectures, err error) {
	prefectures, err = interactor.PrefectureRepository.FindAll()
	return
}

func (interactor *PrefectureInteractor) Add(p model.Prefecture) (prefecture model.Prefecture, err error) {
	prefecture, err = interactor.PrefectureRepository.Store(p)
	return
}

func (interactor *PrefectureInteractor) Update(p model.Prefecture) (prefecture model.Prefecture, err error) {
	prefecture, err = interactor.PrefectureRepository.Update(p)
	return
}

func (interactor *PrefectureInteractor) DeleteById(p model.Prefecture) (err error) {
	err = interactor.PrefectureRepository.DeleteById(p)
	return
}
