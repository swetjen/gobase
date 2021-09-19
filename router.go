package main

import (
	"github.com/labstack/echo/v4"
	"monkey/handlers"
)

func MountApiRouter(e *echo.Echo) {
	// JWT config
	// Router with API Prefix
	// Issues
	issues := e.Group("/issues")
	issues.GET("/", handlers.IssuesGetAll)
	issues.POST("/", handlers.IssueCreate)
	issues.GET("/id/:id", handlers.IssueGetOneById)
	issues.GET("/code/:code", handlers.IssueGetOneByCode)

	// Contacts
	contacts := e.Group("/contacts")
	contacts.POST("/", handlers.ContactsCreate)
	contacts.GET("/", handlers.ContactsGetAll)

	// Login
	e.POST("/login/access-token", handlers.LoginAccessToken)

	// Admin routes
	admin := e.Group("/admin")
	// Admin Users
	admin.GET("/users", handlers.UsersGetAll)
}
