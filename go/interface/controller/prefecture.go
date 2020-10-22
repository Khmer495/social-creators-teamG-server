package controller

import (
	"strconv"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
	"github.com/Khmer495/social-creators-teamG-server/go/usecase"
	"github.com/labstack/echo/v4"
)

type PrefectureController struct {
	Interactor usecase.PrefectureInteractor
}

func NewPrefectureController(sqlHandler database.SqlHandler) *PrefectureController {
	return &PrefectureController{
		Interactor: usecase.PrefectureInteractor{
			PrefectureRepository: &database.PrefectureRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PrefectureController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	prefecture, err := controller.Interactor.PrefectureById(id)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, prefecture)
	return
}

func (controller *PrefectureController) Index(c echo.Context) (err error) {
	prefectures, err := controller.Interactor.Prefectures()
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, prefectures)
	return
}

func (controller *PrefectureController) Create(c echo.Context) (err error) {
	p := model.Prefecture{}
	_ = c.Bind(&p)
	prefecture, err := controller.Interactor.Add(p)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(201, prefecture)
	return
}

func (controller *PrefectureController) Save(c echo.Context) (err error) {
	p := model.Prefecture{}
	_ = c.Bind(&p)
	prefecture, err := controller.Interactor.Update(p)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, prefecture)
	return
}

func (controller *PrefectureController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	prefecture := model.Prefecture{
		ID: id,
	}
	err = controller.Interactor.DeleteById(prefecture)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(204, nil)
	return
}
