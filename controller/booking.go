package controller

import (
	"parkezy/server/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type BookParkingInput struct {
	UserId      int    `json:"user_id" binding:"required"`
	ParkingName string `json:"parking_name" binding:"required"`
	Model       string `json:"model" binding:"required"`
	Plate       string `json:"plate" binding:"required"`
}

type BookingResult struct {
	Id        int    `json:"id"`
	Model     string `json:"model"`
	Plate     string `json:"plate"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func BookParkingSlot(c *gin.Context) {
	var input BookParkingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// get parking
	var parking model.Parking
	if err := model.DB.Where("name = ?", input.ParkingName).First(&parking).Error; err != nil {
		c.JSON(400, gin.H{"error": "Parking not found"})
		return
	}

	upperPlate := strings.ToUpper(input.Plate)
	// create booking
	booking := model.Booking{UserID: input.UserId, ParkingID: int(parking.ID), Model: input.Model, Plate: upperPlate}
	model.DB.Create(&booking)
	c.JSON(200, gin.H{"data": booking})
}

// Get the bookings related to a parking place
func GetBookings(c *gin.Context) {
	var booking []BookingResult
	parkingId := c.Param("parkingId")
	// convert out to int
	out := c.Query("out")
	var outInt int
	if out == "true" || out == "1" {
		outInt = 1
	} else {
		outInt = 0
	}

	// get parking id from parking name
	var parking model.Parking
	if err := model.DB.Where("name = ?", parkingId).First(&parking).Error; err != nil {
		c.JSON(400, gin.H{"error": "Parking not found"})
		return
	}

	// get all bookings
	model.DB.Table("bookings").Where("parking_id = ? and bookings.`out` = ?", parking.ID, outInt).Select("id, model, plate, created_at, updated_at").Scan(&booking)
	c.JSON(200, gin.H{"data": booking})
}

type UpdateBookingInput struct {
	ID int `json:"id" binding:"required"`
}

func UpdateBooking(c *gin.Context) {
	var input UpdateBookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// update booking where id = input.id
	model.DB.Model(&model.Booking{}).Where("id = ?", input.ID).Update("out", 1)
}
