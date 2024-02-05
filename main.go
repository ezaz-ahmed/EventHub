package main

import (
	"net/http"

	"github.com/ezaz-ahmed/EventHub/db"
	"github.com/ezaz-ahmed/EventHub/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InintDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cound not fetch events. Try again later!"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserId = 39
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created Successfully!", "event": event})
}
