package main

import (
	"alexpereap/pereaperformance-backend.git/controllers"
	"alexpereap/pereaperformance-backend.git/database"
	"alexpereap/pereaperformance-backend.git/middlewares"
	"alexpereap/pereaperformance-backend.git/repository"
	"alexpereap/pereaperformance-backend.git/routes"
	"alexpereap/pereaperformance-backend.git/service"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB
	dbConn, err := database.Connect()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// DI / wiring
	userRepository := repository.NewUserRepository(dbConn)
	slideRepository := repository.NewSlideRepository(dbConn)

	userService := service.NewUserService(userRepository)
	loginService := service.NewLoginService(userService)
	slideService := service.NewSlideService(slideRepository)

	cmsController := controllers.NewCmsController()
	userController := controllers.NewUserController(userService)
	loginController := controllers.NewLoginController(loginService)
	slideController := controllers.NewSlideController(slideService)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	if err := setRecursiveTemplates(server); err != nil {
		panic(err)
	}
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
		Slides:       slideController,
		AuthRequired: middlewares.AuthRequired, // middleware factory
	})

	// run
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}

func setRecursiveTemplates(server *gin.Engine) error {
	tpl := template.New("")
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".html") {
			_, perr := tpl.ParseFiles(path)
			return perr
		}
		return nil
	})
	if err != nil {
		return err
	}
	server.SetHTMLTemplate(tpl)
	return nil
}
