package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type RestaurantInteractor struct {
	RestaurantRepository RestaurantRepository
}

func (interactor *RestaurantInteractor) RestaurantById(id int) (restaurant model.Restaurant, err error) {
	restaurant, err = interactor.RestaurantRepository.FindById(id)
	return
}

func (interactor *RestaurantInteractor) Restaurants() (restaurants model.Restaurants, err error) {
	restaurants, err = interactor.RestaurantRepository.FindAll()
	return
}

func (interactor *RestaurantInteractor) Add(r model.Restaurant) (restaurant model.Restaurant, err error) {
	restaurant, err = interactor.RestaurantRepository.Store(r)
	return
}

func (interactor *RestaurantInteractor) Update(r model.Restaurant) (restaurant model.Restaurant, err error) {
	restaurant, err = interactor.RestaurantRepository.Update(r)
	return
}

func (interactor *RestaurantInteractor) DeleteById(r model.Restaurant) (err error) {
	err = interactor.RestaurantRepository.DeleteById(r)
	return
}
