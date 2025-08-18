package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CmsController interface {
	Login(ctx *gin.Context)
	Dashboard(ctx *gin.Context)
}

type cmsController struct{}

func NewCmsController() CmsController {
	return &cmsController{}
}

func (c *cmsController) Login(ctx *gin.Context) {
	data := gin.H{
		"title": "Perea Performance - CMS login",
	}

	sess := sessions.Default(ctx)

	if sess.Get("user") != nil {
		ctx.Redirect(http.StatusFound, "/cms/dashboard")
		return
	}

	ctx.HTML(http.StatusOK, "cms/login.html", data)
}

func (c *cmsController) Dashboard(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ok in dashboard"})
}
