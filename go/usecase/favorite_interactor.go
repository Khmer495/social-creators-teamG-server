package usecase

import "github.com/Khmer495/social-creators-teamG-server/go/entity/model"

type FavoriteInteractor struct {
	FavoriteRepository FavoriteRepository
}

func (interactor *FavoriteInteractor) FavoriteById(id int) (favorite model.Favorite, err error) {
	favorite, err = interactor.FavoriteRepository.FindById(id)
	return
}

func (interactor *FavoriteInteractor) Favorites() (favorites model.Favorites, err error) {
	favorites, err = interactor.FavoriteRepository.FindAll()
	return
}

func (interactor *FavoriteInteractor) Add(f model.Favorite) (favorite model.Favorite, err error) {
	favorite, err = interactor.FavoriteRepository.Store(f)
	return
}

func (interactor *FavoriteInteractor) Update(f model.Favorite) (favorite model.Favorite, err error) {
	favorite, err = interactor.FavoriteRepository.Update(f)
	return
}

func (interactor *FavoriteInteractor) DeleteById(f model.Favorite) (err error) {
	err = interactor.FavoriteRepository.DeleteById(f)
	return
}
