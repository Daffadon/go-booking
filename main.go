package main

import (
	"go-booking/cmd"
	"go-booking/config"
	"go-booking/controller"
	"go-booking/middleware"
	"go-booking/repository"
	"go-booking/routes"
	"go-booking/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	defer config.CloseDB(db)

	if len(os.Args) > 1 {
		cmd.DatabaseCommand(db)
	}

	var (
		jwtService = service.NewJWTService()
		//  User Module Dependency Injection
		userRepository = repository.NewUserRepository(db)
		userService    = service.NewUserService(userRepository, jwtService)
		userController = controller.NewUserController(userService)

		// Book Module Dependency Injection
		bookRepository = repository.NewBookRepository(db)
		bookService    = service.NewBookService(bookRepository)
		bookController = controller.NewBookController(bookService)
	)
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.UserRoute(server, userController, jwtService)
	routes.BookRoute(server, bookController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "development" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
