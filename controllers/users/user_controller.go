package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afaf-tech/bookstore_users-api/domain/users"
	"github.com/afaf-tech/bookstore_users-api/services"
	"github.com/afaf-tech/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO handle JSON ERROR
		restErr := errors.NewBadRequestError("invalid json body")
		fmt.Println(err)
		// TODO return bad request to the caller
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// todo User Creation Error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	// fmt.Println(string(bytes))
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id must be a number")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(userId)
	if getErr != nil {
		// todo User Creation Error
		c.JSON(getErr.Status, getErr)
		return
	}
	// fmt.Println(user)
	// fmt.Println(string(bytes))
	c.JSON(http.StatusOK, result)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}
