package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type CityInteractor struct {
	CityRepository CityRepository
}

func (interactor *CityInteractor) CityById(id int) (city model.City, err error) {
	city, err = interactor.CityRepository.FindById(id)
	return
}

func (interactor *CityInteractor) Citys() (citys model.Cities, err error) {
	citys, err = interactor.CityRepository.FindAll()
	return
}

func (interactor *CityInteractor) Add(c model.City) (city model.City, err error) {
	city, err = interactor.CityRepository.Store(c)
	return
}

func (interactor *CityInteractor) Update(c model.City) (city model.City, err error) {
	city, err = interactor.CityRepository.Update(c)
	return
}

func (interactor *CityInteractor) DeleteById(c model.City) (err error) {
	err = interactor.CityRepository.DeleteById(c)
	return
}
