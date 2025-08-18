package controllers

import (
	"net/http"

	"alexpereap/pereaperformance-backend.git/service"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginController interface {
	LoginHandler(c *gin.Context)
}

type loginController struct {
	service service.LoginService
}

func NewLoginController(service service.LoginService) LoginController {
	return &loginController{
		service: service,
	}
}

func (l *loginController) LoginHandler(c *gin.Context) {
	var form LoginForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username & password are required"})
		return
	}

	err := l.service.LogInUser(form.Username, form.Password, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/cms/dashboard")
}
