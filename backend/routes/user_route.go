package routes

import (
    "firstapp/controllers"
    "github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
    e.POST("/user", controllers.CreateUser)
    e.GET("/user/:userId", controllers.GetAUser)
    e.PUT("/user/:userId", controllers.EditAUser)
    e.DELETE("/user/:userId", controllers.DeleteAUser)
    e.GET("/users", controllers.GetAllUsers)
}