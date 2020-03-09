package app

import (
	"github.com/devMiguelFerrer/book_users-api/controllers/ping"
	"github.com/devMiguelFerrer/book_users-api/controllers/users"
)

func urlMapping() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.GetUser)
	router.GET("/users", users.SearchUser)
}
