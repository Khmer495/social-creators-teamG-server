package controller

import (
	"strconv"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
	"github.com/Khmer495/social-creators-teamG-server/go/usecase"
	"github.com/labstack/echo/v4"
)

type CityController struct {
	Interactor usecase.CityInteractor
}

func NewCityController(sqlHandler database.SqlHandler) *CityController {
	return &CityController{
		Interactor: usecase.CityInteractor{
			CityRepository: &database.CityRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CityController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	city, err := controller.Interactor.CityById(id)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, city)
	return
}

func (controller *CityController) Index(c echo.Context) (err error) {
	cities, err := controller.Interactor.Citys()
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, cities)
	return
}

func (controller *CityController) Create(c echo.Context) (err error) {
	t := model.City{}
	_ = c.Bind(&t)
	city, err := controller.Interactor.Add(t)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(201, city)
	return
}

func (controller *CityController) Save(c echo.Context) (err error) {
	t := model.City{}
	_ = c.Bind(&t)
	city, err := controller.Interactor.Update(t)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, city)
	return
}

func (controller *CityController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	city := model.City{
		ID: id,
	}
	err = controller.Interactor.DeleteById(city)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(204, nil)
	return
}
