package app

import (
	"github.com/afaf-tech/bookstore_users-api/controllers/ping"
	"github.com/afaf-tech/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
