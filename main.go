package main

import (
	"alexpereap/pereaperformance-backend.git/controllers"
	"alexpereap/pereaperformance-backend.git/middlewares"
	"alexpereap/pereaperformance-backend.git/repository"
	"alexpereap/pereaperformance-backend.git/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService    service.UserService       = service.NewUserService(userRepository)
	loginService   service.LoginService      = service.NewLoginService(userService)

	cmsController   controllers.CmsController   = controllers.NewCmsController()
	userController  controllers.UserController  = controllers.NewUserController(userService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	server.LoadHTMLGlob("templates/**/*.html")
	server.Static("/css", "./assets/css")

	secret := []byte("FLrhYZvn4ccpecNZ2jlGg5VIFTscEy4O")
	store := cookie.NewStore(secret)
	// Opciones de seguridad de la cookie (ajusta según tu entorno)
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   int((24 * time.Hour).Seconds()),
		HttpOnly: true,
		Secure:   false, // ponlo en true si sirves por HTTPS
		SameSite: http.SameSiteLaxMode,
	})
	server.Use(sessions.Sessions("sid", store))

	cmsRoutes := server.Group("/cms", middlewares.AuthRequired())
	{
		cmsRoutes.GET("/", func(ctx *gin.Context) {
			sess := sessions.Default(ctx)
			if sess.Get("user") == nil {
				ctx.Redirect(http.StatusFound, "/cms/login")
				return
			}

			ctx.Redirect(http.StatusFound, "/cms/dashboard")
		})

		cmsRoutes.GET("/login", cmsController.Login)
		cmsRoutes.GET("/dashboard", cmsController.Dashboard)
		cmsRoutes.POST("/login", func(ctx *gin.Context) {
			loginController.LoginHandler(ctx)
		})

	}

	userRoutes := server.Group("/users")
	{
		userRoutes.GET("/", func(ctx *gin.Context) {
			users := userController.FindAll()
			ctx.JSON(200, users)

		})

		userRoutes.POST("/", func(ctx *gin.Context) {
			user, err := userController.Save(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, user)
			}
		})

		userRoutes.DELETE("/:id", func(ctx *gin.Context) {
			err := userController.Delete(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			}
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
