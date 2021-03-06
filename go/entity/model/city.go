package model

import "time"

type City struct {
	ID           int        `gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
	PrefectureID int        `json:"prefecture_id"`
	Prefecture   Prefecture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"prefecture"`
	Name         string     `json:"name"`
	Order        int        `json:"order"`
}

type Cities []City
