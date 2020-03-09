package users

import (
	"fmt"

	"github.com/devMiguelFerrer/book_users-api/domain/users"
	"github.com/devMiguelFerrer/book_users-api/services"
	"github.com/devMiguelFerrer/book_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(200, result)
}

func GetUser(c *gin.Context) {

}

func SearchUser(c *gin.Context) {

}
