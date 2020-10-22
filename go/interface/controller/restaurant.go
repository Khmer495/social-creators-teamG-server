package controller

import (
	"log"
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
	log.Println(c)
	prefecture_id, _ := strconv.Atoi(c.QueryParam("prefecture_id"))
	city_id, _ := strconv.Atoi(c.QueryParam("city_id"))
	from_open_time := c.QueryParam("from_open_time")
	to_close_time := c.QueryParam("to_close_time")
	from_kids_welcome_time := c.QueryParam("from_kids_welcome_time")
	to_kids_welcome_time := c.QueryParam("to_kids_welcome_time")
	is_ok_baby_car, _ := strconv.ParseBool(c.QueryParam("is_ok_baby_car"))
	is_zashiki, _ := strconv.ParseBool(c.QueryParam("is_zashiki"))
	is_private_room, _ := strconv.ParseBool(c.QueryParam("is_private_room"))
	is_parking_area, _ := strconv.ParseBool(c.QueryParam("is_parking_area"))
	is_luggage_storage, _ := strconv.ParseBool(c.QueryParam("is_luggage_storage"))
	is_nap_space, _ := strconv.ParseBool(c.QueryParam("is_nap_space"))
	is_baby_bed, _ := strconv.ParseBool(c.QueryParam("is_baby_bed"))
	is_kids_menu, _ := strconv.ParseBool(c.QueryParam("is_kids_menu"))
	is_kids_chair, _ := strconv.ParseBool(c.QueryParam("is_kids_chair"))
	is_kids_dish, _ := strconv.ParseBool(c.QueryParam("is_kids_dish"))
	is_kids_toilet, _ := strconv.ParseBool(c.QueryParam("is_kids_toilet"))
	is_kids_toy, _ := strconv.ParseBool(c.QueryParam("is_kids_toy"))
	is_ok_baby_food, _ := strconv.ParseBool(c.QueryParam("is_ok_baby_food"))
	is_nursing_room, _ := strconv.ParseBool(c.QueryParam("is_nursing_room"))
	is_kids_room, _ := strconv.ParseBool(c.QueryParam("is_kids_room"))

	r := model.RestaurantQuery{
		PrefectureID:        prefecture_id,
		CityID:              city_id,
		FromOpenTime:        from_open_time,
		ToCloseTime:         to_close_time,
		FromKidsWelcomeTime: from_kids_welcome_time,
		ToKidsWelcomeTime:   to_kids_welcome_time,
		IsOkBabyCar:         is_ok_baby_car,
		IsZashiki:           is_zashiki,
		IsPrivateRoom:       is_private_room,
		IsParkingArea:       is_parking_area,
		IsLuggageStorage:    is_luggage_storage,
		IsNapSpace:          is_nap_space,
		IsBabyBed:           is_baby_bed,
		IsKidsMenu:          is_kids_menu,
		IsKidsChair:         is_kids_chair,
		IsKidsDish:          is_kids_dish,
		IsKidsToilet:        is_kids_toilet,
		IsKidsToy:           is_kids_toy,
		IsOkBabyFood:        is_ok_baby_food,
		IsNursingRoom:       is_nursing_room,
		IsKidsRoom:          is_kids_room,
	}
	log.Println(&r)
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
