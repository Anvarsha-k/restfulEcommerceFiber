package models

import "time"

type User struct {

	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	Lastname string `json:"last_name"`
}
