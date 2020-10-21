package model

import "time"

type Restaurant struct {
	ID                  int        `gorm:"primary_key"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `sql:"index" json:"deleted_at"`
	Name                string     `json:"name"`
	Address             string     `json:"address"`
	Lat                 float32    `json:"lat"`
	Lng                 float32    `json:"lng"`
	PrefectureID        int        `json:"prefecture_id"`
	Prefecture          Prefecture
	CityID              int `json:"city_id"`
	City                City
	OpenTime            time.Time `json:"open_time"`
	CloseTime           time.Time `json:"close_time"`
	FromKidsWelcomeTime time.Time `json:"from_kids_welcome_time"`
	ToKidsWelcomeTime   time.Time `json:"to_kids_welcome_time"`
	Menu                string    `json:"menu"`
	IsOkBabyCar         bool      `json:"is_ok_baby_car"`
	IsZashiki           bool      `json:"is_zashiki"`
	IsPrivateRoom       bool      `json:"is_private_room"`
	IsParkingArea       bool      `json:"is_parking_area"`
	IsLuggageStorage    bool      `json:"is_luggage_storage"`
	IsNapSpace          bool      `json:"is_nap_space"`
	IsBabyBed           bool      `json:"is_baby_bed"`
	IsKidsMenu          bool      `json:"is_kids_menu"`
	IsKidsChair         bool      `json:"is_kids_chair"`
	IsKidsDish          bool      `json:"is_kids_dish"`
	IsKidsToilet        bool      `json:"is_kids_toilet"`
	IsKidsToy           bool      `json:"is_kids_toy"`
	IsOkBabyFood        bool      `json:"is_ok_baby_food"`
	IsNursingRoom       bool      `json:"is_nursing_room"`
	IsKidsRoom          bool      `json:"is_kids_room"`
}

type Restaurants []Restaurant
