package controller

import (
	"strconv"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
	"github.com/Khmer495/social-creators-teamG-server/go/usecase"
	"github.com/labstack/echo/v4"
)

type FavoriteController struct {
	Interactor usecase.FavoriteInteractor
}

func NewFavoriteController(sqlHandler database.SqlHandler) *FavoriteController {
	return &FavoriteController{
		Interactor: usecase.FavoriteInteractor{
			FavoriteRepository: &database.FavoriteRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *FavoriteController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	favorite, err := controller.Interactor.FavoriteById(id)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, favorite)
	return
}

func (controller *FavoriteController) Index(c echo.Context) (err error) {
	f := model.FavoriteQuery{}
	_ = c.Bind(&f)
	favorites, err := controller.Interactor.Favorites(f)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, favorites)
	return
}

func (controller *FavoriteController) Create(c echo.Context) (err error) {
	f := model.Favorite{}
	_ = c.Bind(&f)
	favorite, err := controller.Interactor.Add(f)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(201, favorite)
	return
}

func (controller *FavoriteController) Save(c echo.Context) (err error) {
	f := model.Favorite{}
	_ = c.Bind(&f)
	favorite, err := controller.Interactor.Update(f)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, favorite)
	return
}

func (controller *FavoriteController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	favorite := model.Favorite{
		ID: id,
	}
	err = controller.Interactor.DeleteById(favorite)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(204, nil)
	return
}
