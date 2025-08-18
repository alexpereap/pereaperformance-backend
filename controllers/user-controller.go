package controllers

import (
	"alexpereap/pereaperformance-backend.git/entity"
	"alexpereap/pereaperformance-backend.git/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Save(ctx *gin.Context) (*entity.User, error)
	Delete(ctx *gin.Context) error
	FindAll() []entity.User
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Save(ctx *gin.Context) (*entity.User, error) {
	var user *entity.User
	err := ctx.ShouldBind(&user)

	if err != nil {
		return user, err
	}

	user = c.service.Save(user)

	return user, nil
}
func (c *userController) Delete(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}

	user.ID = id
	err = c.service.Delete(user)

	if err != nil {
		return err
	}

	return nil
}

func (c *userController) FindAll() []entity.User {
	return c.service.FindAll()
}
