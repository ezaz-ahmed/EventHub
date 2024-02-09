package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events", createEvent)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", func(ctx *gin.Context) {})
}
