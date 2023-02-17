package main

import (
	"parkezy/server/controller"
	"parkezy/server/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	model.ConnectDB()

	// auth
	r.POST("/user", controller.CreateUser)
	r.POST("/login", controller.LoginUser)

	// parking
	r.GET("/parking", controller.GetParkings)

	// book parking
	r.POST("/book", controller.BookParkingSlot)

	// get bookings of a specific parking
	r.GET("/book/:parkingId", controller.GetBookings)

	// update booking set out = 1
	r.POST("/book/updateBooking", controller.UpdateBooking)

	r.Run()
}
