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
	PrefectureID        int        `json:"-"`
	Prefecture          Prefecture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"prefecture"`
	CityID              int        `json:"-"`
	City                City       `json:"city"`
	OpenTime            string     `json:"open_time"`
	CloseTime           string     `json:"close_time"`
	FromKidsWelcomeTime string     `json:"from_kids_welcome_time"`
	ToKidsWelcomeTime   string     `json:"to_kids_welcome_time"`
	Menu                string     `json:"menu"`
	IsOkBabyCar         bool       `json:"is_ok_baby_car"`
	IsZashiki           bool       `json:"is_zashiki"`
	IsPrivateRoom       bool       `json:"is_private_room"`
	IsParkingArea       bool       `json:"is_parking_area"`
	IsLuggageStorage    bool       `json:"is_luggage_storage"`
	IsNapSpace          bool       `json:"is_nap_space"`
	IsBabyBed           bool       `json:"is_baby_bed"`
	IsKidsMenu          bool       `json:"is_kids_menu"`
	IsKidsChair         bool       `json:"is_kids_chair"`
	IsKidsDish          bool       `json:"is_kids_dish"`
	IsKidsToilet        bool       `json:"is_kids_toilet"`
	IsKidsToy           bool       `json:"is_kids_toy"`
	IsOkBabyFood        bool       `json:"is_ok_baby_food"`
	IsNursingRoom       bool       `json:"is_nursing_room"`
	IsKidsRoom          bool       `json:"is_kids_room"`
}

type Restaurants []Restaurant
