package database

import (
	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
)

type RestaurantRepository struct {
	SqlHandler
}

func (repo *RestaurantRepository) FindById(id int) (restaurant model.Restaurant, err error) {
	if err = repo.Find(&restaurant, id).Error; err != nil {
		return
	}
	return
}

func (repo *RestaurantRepository) FindAll(r model.RestaurantQuery) (restaurants model.Restaurants, err error) {
	query := repo.Query()
	if r.PrefectureID > 0 {
		query = query.Where("prefecture_id = ?", r.PrefectureID)
	}
	if r.CityID > 0 {
		query = query.Where("city_id = ?", r.CityID)
	}
	// if r.FromOpenTime != "" {
	// 	query = query.Where("open_time < ?", r.FromOpenTime)
	// }
	// if r.ToCloseTime != "" {
	// 	query = query.Where("to_time > ?", &r.ToCloseTime)
	// }
	// if r.FromKidsWelcomeTime != "" {
	// 	query = query.Where("from_kids_welcome_time < ?", &r.FromKidsWelcomeTime)
	// }
	// if r.ToKidsWelcomeTime != "" {
	// 	query = query.Where("to_kids_welcome_time > ?", &r.ToKidsWelcomeTime)
	// }
	if r.IsOkBaby {
		query = query.Where("is_ok_baby = ?", r.IsOkBaby)
	}
	if r.IsOkInfant {
		query = query.Where("is_ok_infant = ?", r.IsOkInfant)
	}
	if r.IsNoSmoking {
		query = query.Where("is_no_smoking = ?", r.IsNoSmoking)
	}
	if r.IsOkBabyCar {
		query = query.Where("is_ok_baby_car = ?", r.IsOkBabyCar)
	}
	if r.IsZashiki {
		query = query.Where("is_zashiki = ?", r.IsZashiki)
	}
	if r.IsPrivateRoom {
		query = query.Where("is_private_room = ?", &r.IsPrivateRoom)
	}
	if r.IsParkingArea {
		query = query.Where("is_parking_area = ?", &r.IsParkingArea)
	}
	if r.IsLuggageStorage {
		query = query.Where("is_luggage_storage = ?", &r.IsLuggageStorage)
	}
	if r.IsNapSpace {
		query = query.Where("is_nap_space = ?", &r.IsNapSpace)
	}
	if r.IsBabyBed {
		query = query.Where("is_baby_bed = ?", &r.IsBabyBed)
	}
	if r.IsKidsMenu {
		query = query.Where("is_kids_menu = ?", &r.IsKidsMenu)
	}
	if r.IsKidsChair {
		query = query.Where("is_kids_chair = ?", &r.IsKidsChair)
	}
	if r.IsKidsDish {
		query = query.Where("is_kids_dish = ?", &r.IsKidsDish)
	}
	if r.IsKidsToilet {
		query = query.Where("is_kids_toilet = ?", &r.IsKidsToilet)
	}
	if r.IsKidsToy {
		query = query.Where("is_kids_toy = ?", &r.IsKidsToy)
	}
	if r.IsOkBabyFood {
		query = query.Where("is_ok_baby_food = ?", &r.IsOkBabyFood)
	}
	if r.IsNursingRoom {
		query = query.Where("is_nursing_room = ?", &r.IsNursingRoom)
	}
	if r.IsKidsRoom {
		query = query.Where("is_kids_room = ?", &r.IsKidsRoom)
	}
	if r.IsAllergenLabel {
		query = query.Where("is_allergen_label = ?", &r.IsAllergenLabel)
	}
	if err = query.Find(&restaurants).Error; err != nil {
		return
	}
	return
}

func (repo *RestaurantRepository) Store(r model.Restaurant) (restaurant model.Restaurant, err error) {
	if err = repo.Create(&r).Error; err != nil {
		return
	}
	restaurant = r
	return
}

func (repo *RestaurantRepository) Update(r model.Restaurant) (restaurant model.Restaurant, err error) {
	if err = repo.Save(&r).Error; err != nil {
		return
	}
	restaurant = r
	return
}

func (repo *RestaurantRepository) DeleteById(restaurant model.Restaurant) (err error) {
	if err = repo.Delete(&restaurant).Error; err != nil {
		return
	}
	return
}
