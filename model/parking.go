package model

type Parking struct {
	ID   uint    `json:"id" gorm:"primary_key"`
	Name string  `json:"name"`
	Long float64 `json:"long"`
	Lat  float64 `json:"lat"`
}
