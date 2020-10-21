package model

import "time"

type Prefecture struct {
	ID        int        `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Name      string     `json:"name"`
	Order     int        `json:"order"`
}

type Prefectures []Prefecture
