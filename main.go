package main

import (
	"alexpereap/pereaperformance-backend.git/controllers"
	"alexpereap/pereaperformance-backend.git/middlewares"
	"alexpereap/pereaperformance-backend.git/repository"
	"alexpereap/pereaperformance-backend.git/routes"
	"alexpereap/pereaperformance-backend.git/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// DI / wiring
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	loginService := service.NewLoginService(userService)

	cmsController := controllers.NewCmsController()
	userController := controllers.NewUserController(userService)
	loginController := controllers.NewLoginController(loginService)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	server.LoadHTMLGlob("templates/**/*.html")
	server.Static("/css", "./assets/css")

	// sesiones (global scope)
	secret := []byte("FLrhYZvn4ccpecNZ2jlGg5VIFTscEy4O")
	store := cookie.NewStore(secret)
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   int((24 * time.Hour).Seconds()),
		HttpOnly: true,
		Secure:   false, // true si sirves por HTTPS
		SameSite: http.SameSiteLaxMode,
	})
	server.Use(sessions.Sessions("sid", store))

	// Register routes
	routes.Register(server, routes.Dependencies{
		Cms:          cmsController,
		Users:        userController,
		Login:        loginController,
		AuthRequired: middlewares.AuthRequired, // middleware factory
	})

	// run
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
