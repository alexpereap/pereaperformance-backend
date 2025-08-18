package routes

import (
	"alexpereap/pereaperformance-backend.git/controllers"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	Cms   controllers.CmsController
	Users controllers.UserController
	Login controllers.LoginController

	AuthRequired func() gin.HandlerFunc
}
