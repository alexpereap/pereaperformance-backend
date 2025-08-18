package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUsers(r *gin.Engine, d Dependencies) {
	users := r.Group("/users")
	{
		users.GET("/", func(ctx *gin.Context) {
			data := d.Users.FindAll()
			ctx.JSON(http.StatusOK, data)
		})

		users.POST("/", func(ctx *gin.Context) {
			user, err := d.Users.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, user)
		})

		users.DELETE("/:id", func(ctx *gin.Context) {
			if err := d.Users.Delete(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
		})
	}
}
