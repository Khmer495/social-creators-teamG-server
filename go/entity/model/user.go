package model

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int            `json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  mysql.NullTime `json:"deleted_at"`
	FamilyName string         `json:"family_name"`
	FirstName  string         `json:"first_name"`
	Email      string         `json:"email"`
}
