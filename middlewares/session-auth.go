package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		ignored := map[string]bool{
			"/cms/login": true,
			"/cms/":      true,
		}

		if ignored[c.FullPath()] {
			c.Next()
			return
		}

		sess := sessions.Default(c)

		if sess.Get("user") == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		c.Next()
	}
}
