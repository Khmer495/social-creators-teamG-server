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
	Prefecture          Prefecture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"prefecture"`
	CityID              int        `json:"city_id"`
	City                City       `json:"city"`
	OpenTime            string     `json:"open_time"`
	CloseTime           string     `json:"close_time"`
	FromKidsWelcomeTime string     `json:"from_kids_welcome_time"`
	ToKidsWelcomeTime   string     `json:"to_kids_welcome_time"`
	Menu                string     `json:"menu"`
	IsOkBaby            bool       `json:"is_ok_baby"`
	IsOkInfant          bool       `json:"is_ok_infant"`
	IsNoSmoking         bool       `json:"is_no_smoking"`
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
	IsAllergenLabel     bool       `json:"is_allergen_label"`
}

type Restaurants []Restaurant
type RestaurantQuery struct {
	PrefectureID int `query:"prefecture_id"`
	CityID       int `query:"city_id"`
	// FromOpenTime        string `query:"from_open_time"`
	// ToCloseTime         string `query:"to_close_time"`
	// FromKidsWelcomeTime string `query:"from_kids_welcome_time"`
	// ToKidsWelcomeTime   string `query:"to_kids_welcome_time"`
	IsOkBaby         bool `query:"is_ok_baby"`
	IsOkInfant       bool `query:"is_ok_infant"`
	IsNoSmoking      bool `query:"is_no_smoking"`
	IsOkBabyCar      bool `query:"is_ok_baby_car"`
	IsZashiki        bool `query:"is_zashiki"`
	IsPrivateRoom    bool `query:"is_private_room"`
	IsParkingArea    bool `query:"is_parking_area"`
	IsLuggageStorage bool `query:"is_luggage_storage"`
	IsNapSpace       bool `query:"is_nap_space"`
	IsBabyBed        bool `query:"is_baby_bed"`
	IsKidsMenu       bool `query:"is_kids_menu"`
	IsKidsChair      bool `query:"is_kids_chair"`
	IsKidsDish       bool `query:"is_kids_dish"`
	IsKidsToilet     bool `query:"is_kids_toilet"`
	IsKidsToy        bool `query:"is_kids_toy"`
	IsOkBabyFood     bool `query:"is_ok_baby_food"`
	IsNursingRoom    bool `query:"is_nursing_room"`
	IsKidsRoom       bool `query:"is_kids_room"`
	IsAllergenLabel  bool `query:"is_allergen_label"`
}
