package model

import "time"

type City struct {
	ID           int        `gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
	PrefectureID int        `json:"prerfecture_id"`
	Prefecture   Prefecture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name         string     `json:"name"`
	Order        int        `json:"order"`
}

type Cities []City
