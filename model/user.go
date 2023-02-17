package model

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	// add a unique email gorm tag
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
