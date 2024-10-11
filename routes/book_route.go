package routes

import (
	"go-booking/controller"
	"go-booking/middleware"
	"go-booking/service"

	"github.com/gin-gonic/gin"
)

func BookRoute(route *gin.Engine, bookController controller.BookController, jwtService service.JWTService) {
	routes := route.Group("/api/book")
	{
		routes.GET("/book", middleware.AuthMiddleware(jwtService), bookController.GetBooks)
	}
}
