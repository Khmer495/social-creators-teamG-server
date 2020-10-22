package controller

import (
	"strconv"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
	"github.com/Khmer495/social-creators-teamG-server/go/usecase"
	"github.com/labstack/echo/v4"
)

type CommentController struct {
	Interactor usecase.CommentInteractor
}

func NewCommentController(sqlHandler database.SqlHandler) *CommentController {
	return &CommentController{
		Interactor: usecase.CommentInteractor{
			CommentRepository: &database.CommentRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CommentController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment, err := controller.Interactor.CommentById(id)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, comment)
	return
}

func (controller *CommentController) Index(c echo.Context) (err error) {
	user_id, _ := strconv.Atoi(c.QueryParam("user_id"))
	restaurant_id, _ := strconv.Atoi(c.QueryParam("restaurant_id"))
	m := model.CommentQuery{
		UserID:       user_id,
		RestaurantID: restaurant_id,
	}
	comments, err := controller.Interactor.Comments(m)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, comments)
	return
}

func (controller *CommentController) Create(c echo.Context) (err error) {
	m := model.Comment{}
	_ = c.Bind(&m)
	comment, err := controller.Interactor.Add(m)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(201, comment)
	return
}

func (controller *CommentController) Save(c echo.Context) (err error) {
	m := model.Comment{}
	_ = c.Bind(&m)
	comment, err := controller.Interactor.Update(m)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(200, comment)
	return
}

func (controller *CommentController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment := model.Comment{
		ID: id,
	}
	err = controller.Interactor.DeleteById(comment)
	if err != nil {
		_ = c.JSON(500, NewError(err))
		return
	}
	_ = c.JSON(204, nil)
	return
}
