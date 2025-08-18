package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func registerCMS(r *gin.Engine, d Dependencies) {
	cms := r.Group("/cms", d.AuthRequired())
	{
		cms.GET("/", func(ctx *gin.Context) {
			sess := sessions.Default(ctx)
			if sess.Get("user") == nil {
				ctx.Redirect(http.StatusFound, "/cms/login")
				return
			}
			ctx.Redirect(http.StatusFound, "/cms/dashboard")
		})

		cms.GET("/login", d.Cms.Login)
		cms.GET("/dashboard", d.Cms.Dashboard)

		cms.POST("/login", d.Login.LoginHandler)
	}
}
