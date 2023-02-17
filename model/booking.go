package model

import "time"

type Booking struct {
	ID uint `json:"id" gorm:"primary_key"`

	UserID int `json:"user_id"`
	User   User

	ParkingID int `json:"parking_id"`
	Parking   Parking

	Model string `json:"model"`
	Plate string `json:"plate"`
	Out   bool   `json:"out" gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
	// add createdAt && updated at

}
