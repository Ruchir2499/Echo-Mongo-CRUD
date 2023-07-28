package routes

import (
	"echo-mongo-crud/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	//All routes related to users comes here

	//! create
	e.POST("/user", controllers.CreateUser)

	//! get
	e.GET("/user/:userId", controllers.GetAUser)

	//! update
	e.PUT("/user/:userId", controllers.EditAUser)

	//! delete
	e.DELETE("/user/:userId", controllers.DeleteAUser)

	//! all users
	e.GET("/users", controllers.GetAllUsers)

}
