package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type GoodInteractor struct {
	GoodRepository GoodRepository
}

func (interactor *GoodInteractor) GoodById(id int) (good model.Good, err error) {
	good, err = interactor.GoodRepository.FindById(id)
	return
}

func (interactor *GoodInteractor) Goods(g model.GoodQuery) (goods model.Goods, err error) {
	goods, err = interactor.GoodRepository.FindAll(g)
	return
}

func (interactor *GoodInteractor) Add(g model.Good) (good model.Good, err error) {
	good, err = interactor.GoodRepository.Store(g)
	return
}

func (interactor *GoodInteractor) Update(g model.Good) (good model.Good, err error) {
	good, err = interactor.GoodRepository.Update(g)
	return
}

func (interactor *GoodInteractor) DeleteById(g model.Good) (err error) {
	err = interactor.GoodRepository.DeleteById(g)
	return
}
