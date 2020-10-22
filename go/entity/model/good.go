package model

import "time"

type Good struct {
	ID           int        `gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
	UserID       int        `json:"user_id"`
	User         User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:Cascade;" json:"user"`
	RestaurantID int        `json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:Cascade;" json:"restaurant"`
}

type Goods []Good

type GoodQuery struct {
	UserID       int `query:"user_id"`
	RestaurantID int `query:"restaurant_id"`
}
