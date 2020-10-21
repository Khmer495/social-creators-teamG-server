package model

import "time"

type User struct {
	ID         int        `gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
	FamilyName string     `json:"family_name"`
	FirstName  string     `json:"first_name"`
	Email      string     `json:"email"`
}

type Users []User
