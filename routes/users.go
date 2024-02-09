package routes

import (
	"fmt"
	"net/http"

	"github.com/ezaz-ahmed/EventHub/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Signup is failed. Try again later!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully!"})
}
