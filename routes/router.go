package routes

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine, d Dependencies) {
	registerCMS(r, d)
	registerUsers(r, d)
}
