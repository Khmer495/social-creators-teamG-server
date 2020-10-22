package model

import "time"

type Favorite struct {
	ID           int        `gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
	UserID       int        `json:"-"`
	User         User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:Cascade;" json:"user"`
	RestaurantID int        `json:"-"`
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:Cascade;" json:"restaurant"`
}

type Favorites []Favorite
