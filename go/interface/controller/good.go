package controller

import (
	"strconv"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
	"github.com/Khmer495/social-creators-teamG-server/go/usecase"
	"github.com/labstack/echo/v4"
)

type GoodController struct {
	Interactor usecase.GoodInteractor
}

func NewGoodController(sqlHandler database.SqlHandler) *GoodController {
	return &GoodController{
		Interactor: usecase.GoodInteractor{
			GoodRepository: &database.GoodRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *GoodController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	good, err := controller.Interactor.GoodById(id)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, good)
	return
}

func (controller *GoodController) Index(c echo.Context) (err error) {
	user_id, _ := strconv.Atoi(c.QueryParam("user_id"))
	restaurant_id, _ := strconv.Atoi(c.QueryParam("restaurant_id"))
	g := model.GoodQuery{
		UserID:       user_id,
		RestaurantID: restaurant_id,
	}
	goods, err := controller.Interactor.Goods(g)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, goods)
	return
}

func (controller *GoodController) Create(c echo.Context) (err error) {
	g := model.Good{}
	_ = c.Bind(&g)
	good, err := controller.Interactor.Add(g)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(201, good)
	return
}

func (controller *GoodController) Save(c echo.Context) (err error) {
	g := model.Good{}
	_ = c.Bind(&g)
	good, err := controller.Interactor.Update(g)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, good)
	return
}

func (controller *GoodController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	good := model.Good{
		ID: id,
	}
	err = controller.Interactor.DeleteById(good)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(204, nil)
	return
}
