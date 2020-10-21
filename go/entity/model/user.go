package model

import "time"

type User struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	FamilyName string     `json:"family_name"`
	FirstName  string     `json:"first_name"`
	Email      string     `json:"email"`
}

type Users []User
