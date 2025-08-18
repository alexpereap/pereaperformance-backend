package service

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	LogInUser(username string, password string, ctx *gin.Context) error
	LogOutUser(ctx *gin.Context) error
}

type loginService struct {
	userService UserService
}

func NewLoginService(userService UserService) LoginService {
	return &loginService{
		userService: userService,
	}
}

func (l *loginService) LogInUser(username string, password string, ctx *gin.Context) error {

	user := l.userService.FindOne(map[string]interface{}{
		"username": username,
	})

	if user.ID == 0 {
		return errors.New("user credentials invalid")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("user credentials invalid")
	}

	sess := sessions.Default(ctx)
	sess.Set("user", user.Username)

	if err := sess.Save(); err != nil {
		return errors.New("Error creating session " + err.Error())
	}

	return nil
}

func (l *loginService) LogOutUser(ctx *gin.Context) error {
	sess := sessions.Default(ctx)
	sess.Clear()

	if err := sess.Save(); err != nil {
		return errors.New("Error clearing session " + err.Error())
	}

	return nil
}
