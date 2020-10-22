package controller

import (
	"strconv"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
	"github.com/Khmer495/social-creators-teamG-server/go/usecase"
	"github.com/labstack/echo/v4"
)

type RestaurantController struct {
	Interactor usecase.RestaurantInteractor
}

func NewRestaurantController(sqlHandler database.SqlHandler) *RestaurantController {
	return &RestaurantController{
		Interactor: usecase.RestaurantInteractor{
			RestaurantRepository: &database.RestaurantRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RestaurantController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	restaurant, err := controller.Interactor.RestaurantById(id)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, restaurant)
	return
}

func (controller *RestaurantController) Index(c echo.Context) (err error) {
	r := model.RestaurantQuery{}
	_ = c.Bind(&r)
	restaurants, err := controller.Interactor.Restaurants(r)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, restaurants)
	return
}

func (controller *RestaurantController) Create(c echo.Context) (err error) {
	r := model.Restaurant{}
	_ = c.Bind(&r)
	restaurant, err := controller.Interactor.Add(r)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(201, restaurant)
	return
}

func (controller *RestaurantController) Save(c echo.Context) (err error) {
	r := model.Restaurant{}
	_ = c.Bind(&r)
	restaurant, err := controller.Interactor.Update(r)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, restaurant)
	return
}

func (controller *RestaurantController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	restaurant := model.Restaurant{
		ID: id,
	}
	err = controller.Interactor.DeleteById(restaurant)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(204, nil)
	return
}
