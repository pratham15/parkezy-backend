package controller

import (
	"parkezy/server/model"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Name     string `json:"name"`
	Bookings int    `json:"bookings"`
}

/*
Get all parking places adjoined with the number of bookings
*/
func GetParkings(c *gin.Context) {

	var result []Result
	model.DB.Table("parkings").Select("parkings.name as name, count(b.id) as bookings").Joins("left join bookings b on parkings.id = b.parking_id").Group("parkings.name").Order("name ASC").Scan(&result)
	c.JSON(200, gin.H{"data": result})
}
