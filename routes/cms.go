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
		cms.GET("/logout", d.Login.LogOutHandler)
		cms.GET("/dashboard", d.Cms.Dashboard)

		cms.GET("/slides", d.Slides.Dashboard)
		cms.GET("/slides/create", d.Slides.CreateForm)
		cms.POST("/slides/create", func(ctx *gin.Context) {
			_, err := d.Slides.Save(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			sessions := sessions.Default(ctx)
			sessions.Set("flash-success", "Slide created")
			sessions.Save()
			ctx.Redirect(http.StatusFound, "/cms/dashboard")
		})

		cms.POST("/login", d.Login.LoginHandler)
	}
}
